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
	"fmt"

	"github.com/azadeveloper18/tgramm/marmota/pkg/stores/sqlx"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/service/biz/user/internal/dal/dataobject"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserContactsDAO struct {
	db *sqlx.DB
}

func NewUserContactsDAO(db *sqlx.DB) *UserContactsDAO {
	return &UserContactsDAO{
		db: db,
	}
}

// InsertOrUpdate
// insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on conflict (owner_user_id, contact_user_id) do update set contact_phone = EXCLUDED.contact_phone, contact_first_name = EXCLUDED.contact_first_name, contact_last_name = EXCLUDED.contact_last_name, mutual = EXCLUDED.mutual, date2 = EXCLUDED.date2, is_deleted = false
func (dao *UserContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on conflict (owner_user_id, contact_user_id) do update set contact_phone = EXCLUDED.contact_phone, contact_first_name = EXCLUDED.contact_first_name, contact_last_name = EXCLUDED.contact_last_name, mutual = EXCLUDED.mutual, date2 = EXCLUDED.date2, is_deleted = false"

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
// insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on conflict (owner_user_id, contact_user_id) do update set contact_phone = EXCLUDED.contact_phone, contact_first_name = EXCLUDED.contact_first_name, contact_last_name = EXCLUDED.contact_last_name, mutual = EXCLUDED.mutual, date2 = EXCLUDED.date2, is_deleted = false
func (dao *UserContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query string
		r     sql.Result
	)
	query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on conflict (owner_user_id, contact_user_id) do update set contact_phone = EXCLUDED.contact_phone, contact_first_name = EXCLUDED.contact_first_name, contact_last_name = EXCLUDED.contact_last_name, mutual = EXCLUDED.mutual, date2 = EXCLUDED.date2, is_deleted = false"

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

// SelectContact
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where is_deleted = false and owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
func (dao *UserContactsDAO) SelectContact(ctx context.Context, ownerUserId int64, contactUserId int64) (rValue *dataobject.UserContactsDO, err error) {
	var (
		query string
		do    = &dataobject.UserContactsDO{}
	)
	query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where is_deleted = false and owner_user_id = ? and contact_user_id = ?"

	err = dao.db.QueryRowPartial(ctx, do, query, ownerUserId, contactUserId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectContact(_), error: %v", err)
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

// SelectByContactId
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id and is_deleted = false
func (dao *UserContactsDAO) SelectByContactId(ctx context.Context, ownerUserId int64, contactUserId int64) (rValue *dataobject.UserContactsDO, err error) {
	var (
		query string
		do    = &dataobject.UserContactsDO{}
	)
	query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_user_id = ? and is_deleted = false"

	err = dao.db.QueryRowPartial(ctx, do, query, ownerUserId, contactUserId)

	if err != nil {
		if !errors.Is(err, sqlx.ErrNotFound) {
			logx.WithContext(ctx).Errorf("queryx in SelectByContactId(_), error: %v", err)
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

// SelectListByPhoneList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_phone in (:phoneList) and is_deleted = false
func (dao *UserContactsDAO) SelectListByPhoneList(ctx context.Context, ownerUserId int64, phoneList []string) (rList []dataobject.UserContactsDO, err error) {

	if len(phoneList) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_phone in (%s) and is_deleted = false", sqlx.InStringList(phoneList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhoneList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByPhoneListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_phone in (:phoneList) and is_deleted = false
func (dao *UserContactsDAO) SelectListByPhoneListWithCB(ctx context.Context, ownerUserId int64, phoneList []string, cb func(sz, i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {

	if len(phoneList) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_phone in (%s) and is_deleted = false", sqlx.InStringList(phoneList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByPhoneList(_), error: %v", err)
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

// SelectAllUserContacts
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = false
func (dao *UserContactsDAO) SelectAllUserContacts(ctx context.Context, ownerUserId int64) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = false"

	err = dao.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllUserContacts(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectAllUserContactsWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = false
func (dao *UserContactsDAO) SelectAllUserContactsWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = false"

	err = dao.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectAllUserContacts(_), error: %v", err)
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

// SelectUserContacts
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = false order by contact_user_id asc
func (dao *UserContactsDAO) SelectUserContacts(ctx context.Context, ownerUserId int64) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = false order by contact_user_id asc"

	err = dao.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUserContacts(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectUserContactsWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and is_deleted = false order by contact_user_id asc
func (dao *UserContactsDAO) SelectUserContactsWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {
	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and is_deleted = false order by contact_user_id asc"

	err = dao.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectUserContacts(_), error: %v", err)
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

// SelectUserContactIdList
// select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = false order by contact_user_id asc
func (dao *UserContactsDAO) SelectUserContactIdList(ctx context.Context, ownerUserId int64) (rList []int64, err error) {
	var query string
	query = "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = false order by contact_user_id asc"

	err = dao.db.QueryRowsPartial(ctx, &rList, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserContactIdList(_), error: %v", err)
	}

	return
}

// SelectUserContactIdListWithCB
// select contact_user_id from user_contacts where owner_user_id = :owner_user_id and is_deleted = false order by contact_user_id asc
func (dao *UserContactsDAO) SelectUserContactIdListWithCB(ctx context.Context, ownerUserId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query string
	query = "select contact_user_id from user_contacts where owner_user_id = ? and is_deleted = false order by contact_user_id asc"

	err = dao.db.QueryRowsPartial(ctx, &rList, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserContactIdList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectListByIdList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id in (:id_list) and is_deleted = false order by contact_user_id asc
func (dao *UserContactsDAO) SelectListByIdList(ctx context.Context, ownerUserId int64, idList []int64) (rList []dataobject.UserContactsDO, err error) {

	if len(idList) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_user_id in (%s) and is_deleted = false order by contact_user_id asc", sqlx.InInt64List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByIdListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = :owner_user_id and contact_user_id in (:id_list) and is_deleted = false order by contact_user_id asc
func (dao *UserContactsDAO) SelectListByIdListWithCB(ctx context.Context, ownerUserId int64, idList []int64, cb func(sz, i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {

	if len(idList) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id = ? and contact_user_id in (%s) and is_deleted = false order by contact_user_id asc", sqlx.InInt64List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByIdList(_), error: %v", err)
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

// SelectListByOwnerListAndContactList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id in (:idList1) and contact_user_id in (:idList2) and is_deleted = false
func (dao *UserContactsDAO) SelectListByOwnerListAndContactList(ctx context.Context, idList1 []int64, idList2 []int64) (rList []dataobject.UserContactsDO, err error) {
	if len(idList1) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	if len(idList2) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id in (%s) and contact_user_id in (%s) and is_deleted = false", sqlx.InInt64List(idList1), sqlx.InInt64List(idList2))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByOwnerListAndContactList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectListByOwnerListAndContactListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id in (:idList1) and contact_user_id in (:idList2) and is_deleted = false
func (dao *UserContactsDAO) SelectListByOwnerListAndContactListWithCB(ctx context.Context, idList1 []int64, idList2 []int64, cb func(sz, i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {
	if len(idList1) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	if len(idList2) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where owner_user_id in (%s) and contact_user_id in (%s) and is_deleted = false", sqlx.InInt64List(idList1), sqlx.InInt64List(idList2))

	err = dao.db.QueryRowsPartial(ctx, &values, query)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectListByOwnerListAndContactList(_), error: %v", err)
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

// UpdateContactNameById
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = false where id = :id
func (dao *UserContactsDAO) UpdateContactNameById(ctx context.Context, contactFirstName string, contactLastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = false where id = ?"

	rResult, err = dao.db.Exec(ctx, query, contactFirstName, contactLastName, id)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateContactNameById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateContactNameById(_), error: %v", err)
	}

	return
}

// UpdateContactNameByIdTx
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = false where id = :id
func (dao *UserContactsDAO) UpdateContactNameByIdTx(tx *sqlx.Tx, contactFirstName string, contactLastName string, id int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = false where id = ?"

	rResult, err = tx.Exec(query, contactFirstName, contactLastName, id)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateContactNameById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateContactNameById(_), error: %v", err)
	}

	return
}

// UpdateContactName
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = false where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (dao *UserContactsDAO) UpdateContactName(ctx context.Context, contactFirstName string, contactLastName string, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = false where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"

	rResult, err = dao.db.Exec(ctx, query, contactFirstName, contactLastName, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

// UpdateContactNameTx
// update user_contacts set contact_first_name = :contact_first_name, contact_last_name = :contact_last_name, is_deleted = false where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (dao *UserContactsDAO) UpdateContactNameTx(tx *sqlx.Tx, contactFirstName string, contactLastName string, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = false where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"

	rResult, err = tx.Exec(query, contactFirstName, contactLastName, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

// UpdateMutual
// update user_contacts set mutual = :mutual where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (dao *UserContactsDAO) UpdateMutual(ctx context.Context, mutual bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set mutual = ? where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"

	rResult, err = dao.db.Exec(ctx, query, mutual, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateMutual(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateMutual(_), error: %v", err)
	}

	return
}

// UpdateMutualTx
// update user_contacts set mutual = :mutual where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id = :contact_user_id)
func (dao *UserContactsDAO) UpdateMutualTx(tx *sqlx.Tx, mutual bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set mutual = ? where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"

	rResult, err = tx.Exec(query, mutual, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateMutual(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateMutual(_), error: %v", err)
	}

	return
}

// DeleteContacts
// update user_contacts set is_deleted = true, mutual = false, close_friend = false, stories_hidden = false where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id in (:id_list))
func (dao *UserContactsDAO) DeleteContacts(ctx context.Context, ownerUserId int64, idList []int64) (rowsAffected int64, err error) {

	if len(idList) == 0 {
		return
	}

	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update user_contacts set is_deleted = true, mutual = false, close_friend = false, stories_hidden = false where contact_user_id != 0 and (owner_user_id = ? and contact_user_id in (%s))", sqlx.InInt64List(idList))

	rResult, err = dao.db.Exec(ctx, query, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

// DeleteContactsTx
// update user_contacts set is_deleted = true, mutual = false, close_friend = false, stories_hidden = false where contact_user_id != 0 and (owner_user_id = :owner_user_id and contact_user_id in (:id_list))
func (dao *UserContactsDAO) DeleteContactsTx(tx *sqlx.Tx, ownerUserId int64, idList []int64) (rowsAffected int64, err error) {

	if len(idList) == 0 {
		return
	}
	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update user_contacts set is_deleted = true, mutual = false, close_friend = false, stories_hidden = false where contact_user_id != 0 and (owner_user_id = ? and contact_user_id in (%s))", sqlx.InInt64List(idList))

	rResult, err = tx.Exec(query, ownerUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

// UpdatePhoneByContactId
// update user_contacts set contact_phone = :contact_phone where contact_user_id = :contact_user_id
func (dao *UserContactsDAO) UpdatePhoneByContactId(ctx context.Context, contactPhone string, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set contact_phone = ? where contact_user_id = ?"

	rResult, err = dao.db.Exec(ctx, query, contactPhone, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdatePhoneByContactId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdatePhoneByContactId(_), error: %v", err)
	}

	return
}

// UpdatePhoneByContactIdTx
// update user_contacts set contact_phone = :contact_phone where contact_user_id = :contact_user_id
func (dao *UserContactsDAO) UpdatePhoneByContactIdTx(tx *sqlx.Tx, contactPhone string, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set contact_phone = ? where contact_user_id = ?"

	rResult, err = tx.Exec(query, contactPhone, contactUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdatePhoneByContactId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdatePhoneByContactId(_), error: %v", err)
	}

	return
}

// SelectUserReverseContactIdList
// select owner_user_id from user_contacts where contact_user_id = :contact_user_id and is_deleted = false
func (dao *UserContactsDAO) SelectUserReverseContactIdList(ctx context.Context, contactUserId int64) (rList []int64, err error) {
	var query string
	query = "select owner_user_id from user_contacts where contact_user_id = ? and is_deleted = false"

	err = dao.db.QueryRowsPartial(ctx, &rList, query, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserReverseContactIdList(_), error: %v", err)
	}

	return
}

// SelectUserReverseContactIdListWithCB
// select owner_user_id from user_contacts where contact_user_id = :contact_user_id and is_deleted = false
func (dao *UserContactsDAO) SelectUserReverseContactIdListWithCB(ctx context.Context, contactUserId int64, cb func(sz, i int, v int64)) (rList []int64, err error) {
	var query string
	query = "select owner_user_id from user_contacts where contact_user_id = ? and is_deleted = false"

	err = dao.db.QueryRowsPartial(ctx, &rList, query, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("select in SelectUserReverseContactIdList(_), error: %v", err)
	}

	if cb != nil {
		sz := len(rList)
		for i := range sz {
			cb(sz, i, rList[i])
		}
	}

	return
}

// SelectReverseListByIdList
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = :contact_user_id and owner_user_id in (:id_list) and is_deleted = false
func (dao *UserContactsDAO) SelectReverseListByIdList(ctx context.Context, contactUserId int64, idList []int64) (rList []dataobject.UserContactsDO, err error) {

	if len(idList) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = ? and owner_user_id in (%s) and is_deleted = false", sqlx.InInt64List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectReverseListByIdList(_), error: %v", err)
		return
	}

	rList = values

	return
}

// SelectReverseListByIdListWithCB
// select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = :contact_user_id and owner_user_id in (:id_list) and is_deleted = false
func (dao *UserContactsDAO) SelectReverseListByIdListWithCB(ctx context.Context, contactUserId int64, idList []int64, cb func(sz, i int, v *dataobject.UserContactsDO)) (rList []dataobject.UserContactsDO, err error) {

	if len(idList) == 0 {
		rList = []dataobject.UserContactsDO{}
		return
	}

	var (
		query  string
		values []dataobject.UserContactsDO
	)
	query = fmt.Sprintf("select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, close_friend, stories_hidden, date2, is_deleted from user_contacts where contact_user_id = ? and owner_user_id in (%s) and is_deleted = false", sqlx.InInt64List(idList))

	err = dao.db.QueryRowsPartial(ctx, &values, query, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("queryx in SelectReverseListByIdList(_), error: %v", err)
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

// UpdateCloseFriend
// update user_contacts set close_friend = :close_friend where owner_user_id = :owner_user_id and contact_user_id in (:idList)
func (dao *UserContactsDAO) UpdateCloseFriend(ctx context.Context, closeFriend bool, ownerUserId int64, idList []int64) (rowsAffected int64, err error) {

	if len(idList) == 0 {
		return
	}

	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update user_contacts set close_friend = ? where owner_user_id = ? and contact_user_id in (%s)", sqlx.InInt64List(idList))

	rResult, err = dao.db.Exec(ctx, query, closeFriend, ownerUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateCloseFriend(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateCloseFriend(_), error: %v", err)
	}

	return
}

// UpdateCloseFriendTx
// update user_contacts set close_friend = :close_friend where owner_user_id = :owner_user_id and contact_user_id in (:idList)
func (dao *UserContactsDAO) UpdateCloseFriendTx(tx *sqlx.Tx, closeFriend bool, ownerUserId int64, idList []int64) (rowsAffected int64, err error) {

	if len(idList) == 0 {
		return
	}
	var (
		query   string
		rResult sql.Result
	)
	query = fmt.Sprintf("update user_contacts set close_friend = ? where owner_user_id = ? and contact_user_id in (%s)", sqlx.InInt64List(idList))

	rResult, err = tx.Exec(query, closeFriend, ownerUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateCloseFriend(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateCloseFriend(_), error: %v", err)
	}

	return
}

// UpdateStoriesHidden
// update user_contacts set stories_hidden = :stories_hidden where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
func (dao *UserContactsDAO) UpdateStoriesHidden(ctx context.Context, storiesHidden bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set stories_hidden = ? where owner_user_id = ? and contact_user_id = ?"

	rResult, err = dao.db.Exec(ctx, query, storiesHidden, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(ctx).Errorf("exec in UpdateStoriesHidden(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(ctx).Errorf("rowsAffected in UpdateStoriesHidden(_), error: %v", err)
	}

	return
}

// UpdateStoriesHiddenTx
// update user_contacts set stories_hidden = :stories_hidden where owner_user_id = :owner_user_id and contact_user_id = :contact_user_id
func (dao *UserContactsDAO) UpdateStoriesHiddenTx(tx *sqlx.Tx, storiesHidden bool, ownerUserId int64, contactUserId int64) (rowsAffected int64, err error) {
	var (
		query   string
		rResult sql.Result
	)
	query = "update user_contacts set stories_hidden = ? where owner_user_id = ? and contact_user_id = ?"

	rResult, err = tx.Exec(query, storiesHidden, ownerUserId, contactUserId)

	if err != nil {
		logx.WithContext(tx.Context()).Errorf("exec in UpdateStoriesHidden(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		logx.WithContext(tx.Context()).Errorf("rowsAffected in UpdateStoriesHidden(_), error: %v", err)
	}

	return
}
