// Modified by Kgramm, 2026-10-31. Original work remains under the Apache License 2.0.
//
// PostgreSQL access for the sqlx store. Open used to call go-zero NewMysql.
// Callers still pass '?' placeholders and named ':param' queries; those are
// rebound to $1 before they reach the server. INSERT statements that do not
// already say RETURNING id are read back with RETURNING so LastInsertId works
// without the MySQL session counter.

package sqlx

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	gsqlx "github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	pgUniqueViolation    = "23505"
	pgInvalidCatalogName = "3D000"
)

func openPG(c *Config) (*conn, error) {
	db, err := sql.Open("pgx", c.DSN)
	if err != nil {
		return nil, err
	}
	if c.Active > 0 {
		db.SetMaxOpenConns(c.Active)
	}
	if c.Idle > 0 {
		db.SetMaxIdleConns(c.Idle)
	}
	if c.IdleTimeout != "" {
		if d, err := time.ParseDuration(c.IdleTimeout); err == nil {
			db.SetConnMaxLifetime(d)
		}
	}
	return newConn(gsqlx.NewSqlConnFromDB(db)), nil
}

func rebindPG(query string) string {
	return Rebind(DOLLAR, query)
}

func isInsert(query string) bool {
	q := strings.TrimSpace(query)
	return len(q) >= 6 && strings.EqualFold(q[:6], "insert")
}

// pgResult keeps a successful statement successful when the caller asks for
// an insert id PostgreSQL does not store on the session.
type pgResult struct {
	id    int64
	rows  int64
	hasID bool
	inner sql.Result
}

func (r pgResult) LastInsertId() (int64, error) {
	if r.hasID {
		return r.id, nil
	}
	return 0, nil
}

func (r pgResult) RowsAffected() (int64, error) {
	if r.inner != nil {
		return r.inner.RowsAffected()
	}
	return r.rows, nil
}

// insertReturningID is true when an INSERT should be read with QueryRow so
// LastInsertId returns the identity. Callers that already wrote "RETURNING id"
// are included. INSERT ... ON CONFLICT DO NOTHING yields no row; that stays a
// successful result with no id, matching INSERT IGNORE.
func insertReturningID(query string) (string, bool) {
	if !isInsert(query) {
		return query, false
	}
	q := strings.TrimSpace(query)
	lower := strings.ToLower(q)
	if !strings.Contains(lower, "returning") {
		return q + " RETURNING id", true
	}
	trimmed := strings.TrimSuffix(strings.TrimSpace(lower), ";")
	if strings.HasSuffix(strings.TrimSpace(trimmed), "returning id") {
		return q, true
	}
	return q, false
}

func execAdapted(ctx context.Context, sess gsqlx.Session, query string, args ...interface{}) (sql.Result, error) {
	query = rebindPG(query)
	if q, ok := insertReturningID(query); ok {
		var id int64
		err := sess.QueryRowCtx(ctx, &id, q, args...)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) || errors.Is(err, gsqlx.ErrNotFound) {
				return pgResult{}, nil
			}
			return nil, err
		}
		return pgResult{id: id, rows: 1, hasID: true}, nil
	}
	res, err := sess.ExecCtx(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return pgResult{inner: res}, nil
}

func pgState(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}
