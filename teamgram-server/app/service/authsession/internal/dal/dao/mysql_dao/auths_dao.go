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
	"errors"

	"github.com/azadeveloper18/tgramm/marmota/pkg/stores/sqlx"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/service/authsession/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{
		db: db,
	}
}

// InsertOrUpdateLayer
// insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on conflict (auth_key_id) do update set layer = EXCLUDED.layer, client_ip = EXCLUDED.client_ip, date_active = EXCLUDED.date_active
func (dao *AuthsDAO) InsertOrUpdateLayer(ctx context.Context, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on conflict (auth_key_id) do update set layer = EXCLUDED.layer, client_ip = EXCLUDED.client_ip, date_active = EXCLUDED.date_active"

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		logx.WithContext(ctx).Errorf("namedExec in InsertOrUpdateLayer(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(ctx).Errorf("lastInsertId in InsertOrUpdateLayer(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in InsertOrUpdateLayer(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdateLayerTx
// insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on conflict (auth_key_id) do update set layer = EXCLUDED.layer, client_ip = EXCLUDED.client_ip, date_active = EXCLUDED.date_active
func (dao *AuthsDAO) InsertOrUpdateLayerTx(tx *sqlx.Tx, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auths(auth_key_id, layer, api_id, params, client_ip, date_active) values (:auth_key_id, :layer, 0, 'null', :client_ip, :date_active) on conflict (auth_key_id) do update set layer = EXCLUDED.layer, client_ip = EXCLUDED.client_ip, date_active = EXCLUDED.date_active"

	r, err = tx.NamedExec(query, do)
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("namedExec in InsertOrUpdateLayer(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("lastInsertId in InsertOrUpdateLayer(%v), error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in InsertOrUpdateLayer(%v), error: %v", do, err)
	}

	return
}

// InsertOrUpdate
// insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on conflict (auth_key_id) do update set api_id = EXCLUDED.api_id, device_model = EXCLUDED.device_model, system_version = EXCLUDED.system_version, app_version = EXCLUDED.app_version, system_lang_code = EXCLUDED.system_lang_code, lang_pack = EXCLUDED.lang_pack, lang_code = EXCLUDED.lang_code, proxy = EXCLUDED.proxy, params = EXCLUDED.params, client_ip = EXCLUDED.client_ip, date_active = EXCLUDED.date_active
func (dao *AuthsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on conflict (auth_key_id) do update set api_id = EXCLUDED.api_id, device_model = EXCLUDED.device_model, system_version = EXCLUDED.system_version, app_version = EXCLUDED.app_version, system_lang_code = EXCLUDED.system_lang_code, lang_pack = EXCLUDED.lang_pack, lang_code = EXCLUDED.lang_code, proxy = EXCLUDED.proxy, params = EXCLUDED.params, client_ip = EXCLUDED.client_ip, date_active = EXCLUDED.date_active"

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
// insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on conflict (auth_key_id) do update set api_id = EXCLUDED.api_id, device_model = EXCLUDED.device_model, system_version = EXCLUDED.system_version, app_version = EXCLUDED.app_version, system_lang_code = EXCLUDED.system_lang_code, lang_pack = EXCLUDED.lang_pack, lang_code = EXCLUDED.lang_code, proxy = EXCLUDED.proxy, params = EXCLUDED.params, client_ip = EXCLUDED.client_ip, date_active = EXCLUDED.date_active
func (dao *AuthsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into auths(auth_key_id, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip, date_active) values (:auth_key_id, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip, :date_active) on conflict (auth_key_id) do update set api_id = EXCLUDED.api_id, device_model = EXCLUDED.device_model, system_version = EXCLUDED.system_version, app_version = EXCLUDED.app_version, system_lang_code = EXCLUDED.system_lang_code, lang_pack = EXCLUDED.lang_pack, lang_code = EXCLUDED.lang_code, proxy = EXCLUDED.proxy, params = EXCLUDED.params, client_ip = EXCLUDED.client_ip, date_active = EXCLUDED.date_active"

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

// SelectByAuthKeyId
// select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = :auth_key_id and deleted = false limit 1
func (dao *AuthsDAO) SelectByAuthKeyId(ctx context.Context, authKeyId int64) (rValue *dataobject.AuthsDO, err error) {
	var (
		query string
		do    = &dataobject.AuthsDO{}
	)
	query = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip, date_active from auths where auth_key_id = ? and deleted = false limit 1"

	err = dao.db.QueryRowPartial(ctx, do, query, authKeyId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByAuthKeyId(_), error: %v", err)
			return
		} else {
			// not found not error, return nil, nil
			err = nil
		}
	} else {
		rValue = do
	}

	return
}
