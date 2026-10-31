// Modified by Kgramm, 2026-10-31. Original work remains under the Apache License 2.0.
/*
 * WARNING! All changes made in this file will be lost!
 *   Created from by 'dalgen'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package mysql_dao

import (
	"context"
	"database/sql"

	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneBooksDAO struct {
	db *sqlx.DB
}

func NewPhoneBooksDAO(db *sqlx.DB) *PhoneBooksDAO {
	return &PhoneBooksDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on conflict (auth_key_id, client_id) do update set phone = EXCLUDED.phone, first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name
func (dao *PhoneBooksDAO) InsertOrUpdate(ctx context.Context, do *dataobject.PhoneBooksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on conflict (auth_key_id, client_id) do update set phone = EXCLUDED.phone, first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdateTx
// insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on conflict (auth_key_id, client_id) do update set phone = EXCLUDED.phone, first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name
func (dao *PhoneBooksDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.PhoneBooksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on conflict (auth_key_id, client_id) do update set phone = EXCLUDED.phone, first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdate(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdate(%v), error: %v", do, err)
	}

	return
}
