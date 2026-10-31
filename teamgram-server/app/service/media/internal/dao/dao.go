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
	"github.com/azadeveloper18/tgramm/marmota/pkg/stores/sqlc"
	"github.com/azadeveloper18/tgramm/marmota/pkg/stores/sqlx"
	dfs_client "github.com/azadeveloper18/tgramm/teamgram-server/app/service/dfs/client"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/service/media/internal/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type Dao struct {
	*Mysql
	sqlc.CachedConn
	dfs_client.DfsClient
}

func New(c config.Config) *Dao {
	db := sqlx.NewMySQL(&c.Mysql)
	return &Dao{
		Mysql:      newMysqlDao(db),
		CachedConn: sqlc.NewConn(db, c.Cache),
		DfsClient:  dfs_client.NewDfsClient(zrpc.MustNewClient(c.Dfs)),
	}
}
