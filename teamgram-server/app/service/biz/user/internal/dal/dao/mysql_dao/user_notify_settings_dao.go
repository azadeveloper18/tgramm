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
	"github.com/azadeveloper18/tgramm/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserNotifySettingsDAO struct {
	db *sqlx.DB
}

func NewUserNotifySettingsDAO(db *sqlx.DB) *UserNotifySettingsDAO {
	return &UserNotifySettingsDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on conflict (user_id, peer_type, peer_id) do update set show_previews = EXCLUDED.show_previews, silent = EXCLUDED.silent, mute_until = EXCLUDED.mute_until, sound = EXCLUDED.sound, deleted = false
func (dao *UserNotifySettingsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserNotifySettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on conflict (user_id, peer_type, peer_id) do update set show_previews = EXCLUDED.show_previews, silent = EXCLUDED.silent, mute_until = EXCLUDED.mute_until, sound = EXCLUDED.sound, deleted = false"

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
// insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on conflict (user_id, peer_type, peer_id) do update set show_previews = EXCLUDED.show_previews, silent = EXCLUDED.silent, mute_until = EXCLUDED.mute_until, sound = EXCLUDED.sound, deleted = false
func (dao *UserNotifySettingsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserNotifySettingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into user_notify_settings(user_id, peer_type, peer_id, show_previews, silent, mute_until, sound) values (:user_id, :peer_type, :peer_id, :show_previews, :silent, :mute_until, :sound) on conflict (user_id, peer_type, peer_id) do update set show_previews = EXCLUDED.show_previews, silent = EXCLUDED.silent, mute_until = EXCLUDED.mute_until, sound = EXCLUDED.sound, deleted = false"

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

// SelectAll
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and deleted = false
func (dao *UserNotifySettingsDAO) SelectAll(ctx context.Context, userId int64) (rList []dataobject.UserNotifySettingsDO, err error) {
	var (
		query  string
		values []dataobject.UserNotifySettingsDO
	)
	query = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and deleted = false"

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAll(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAllWithCB
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and deleted = false
func (dao *UserNotifySettingsDAO) SelectAllWithCB(ctx context.Context, userId int64, cb func(sz, i int, v *dataobject.UserNotifySettingsDO)) (rList []dataobject.UserNotifySettingsDO, err error) {
	var (
		query  string
		values []dataobject.UserNotifySettingsDO
	)
	query = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and deleted = false"

	err = dao.db.QueryRowsPartial(ctx, &values, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAll(_), error: %v", err)
		return
	}

	rList = values

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, &rList[i])
		}
	}

	return
}

// Select
// select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = :user_id and peer_type = :peer_type and peer_id = :peer_id and deleted = false
func (dao *UserNotifySettingsDAO) Select(ctx context.Context, userId int64, peerType int32, peerId int64) (rValue *dataobject.UserNotifySettingsDO, err error) {
	var (
		query string
		do    = &dataobject.UserNotifySettingsDO{}
	)
	query = "select id, user_id, peer_type, peer_id, show_previews, silent, mute_until, sound from user_notify_settings where user_id = ? and peer_type = ? and peer_id = ? and deleted = false"

	err = dao.db.QueryRowPartial(ctx, do, query, userId, peerType, peerId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in Select(_), error: %v", err)
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

// DeleteAll
// update user_notify_settings set deleted = true where user_id = :user_id
func (dao *UserNotifySettingsDAO) DeleteAll(ctx context.Context, userId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_notify_settings set deleted = true where user_id = ?"

	rResult, err = dao.db.Exec(ctx, query, userId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteAll(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteAll(_), error: %v", err)
	}

	return
}

// DeleteAllTx
// update user_notify_settings set deleted = true where user_id = :user_id
func (dao *UserNotifySettingsDAO) DeleteAllTx(tx *sqlx.Tx, userId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_notify_settings set deleted = true where user_id = ?"

	rResult, err = tx.Exec(query, userId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteAll(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteAll(_), error: %v", err)
	}

	return
}
