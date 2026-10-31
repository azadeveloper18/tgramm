/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Teamgram Authors.
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package users_helper

import (
	"github.com/azadeveloper18/tgramm/teamgram-server/app/bff/users/internal/config"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/bff/users/internal/plugin"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/bff/users/internal/server/grpc/service"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/bff/users/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config, plugin1 plugin.StoryPlugin, plugin2 plugin.PersonalChannelPlugin, plugin3 plugin.WallpaperPlugin) *service.Service {
	return service.New(svc.NewServiceContext(c, plugin1, plugin2, plugin3))
}
