/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package svc

import (
	"github.com/azadeveloper18/tgramm/teamgram-server/app/service/dfs/internal/config"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/service/dfs/internal/dao"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/service/dfs/internal/ffmpegutil"
)

type ServiceContext struct {
	Config config.Config
	*dao.Dao
	*ffmpegutil.FFmpegUtil
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Dao:        dao.New(c),
		FFmpegUtil: ffmpegutil.NewFFmpegUtil(),
	}
}
