// Modified by Kgramm, 2026-10-31. Original work remains under the Apache License 2.0.

package sqlx

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

func TestRebindPG(t *testing.T) {
	got := rebindPG("select * from users where id = ? and deleted = ?")
	require.Equal(t, "select * from users where id = $1 and deleted = $2", got)
	require.Equal(t, "select 1", rebindPG("select 1"))
}

func TestIsDuplicateAndMissingDb(t *testing.T) {
	require.False(t, IsDuplicate(nil))
	require.False(t, IsMissingDb(nil))
	require.False(t, IsDuplicate(errors.New("duplicate")))

	dup := &pgconn.PgError{Code: pgUniqueViolation}
	missing := &pgconn.PgError{Code: pgInvalidCatalogName}
	require.True(t, IsDuplicate(dup))
	require.True(t, IsDuplicate(fmt.Errorf("wrap: %w", dup)))
	require.False(t, IsDuplicate(missing))
	require.True(t, IsMissingDb(missing))
	require.True(t, IsMissingDb(fmt.Errorf("wrap: %w", missing)))
	require.False(t, IsMissingDb(dup))
}
