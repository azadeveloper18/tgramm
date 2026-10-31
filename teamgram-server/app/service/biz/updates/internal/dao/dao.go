/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package dao

import (
	"github.com/azadeveloper18/tgramm/marmota/pkg/net/rpcx"
	"github.com/azadeveloper18/tgramm/marmota/pkg/stores/sqlx"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/service/biz/updates/internal/config"
	idgen_client "github.com/azadeveloper18/tgramm/teamgram-server/app/service/idgen/client"

	"github.com/zeromicro/go-zero/core/stores/kv"
)

type Dao struct {
	*Mysql
	kv kv.Store
	idgen_client.IDGenClient2
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:        newMysqlDao(db),
		kv:           kv.NewStore(c.KV),
		IDGenClient2: idgen_client.NewIDGenClient2(rpcx.GetCachedRpcClient(c.IdgenClient)),
	}
}
