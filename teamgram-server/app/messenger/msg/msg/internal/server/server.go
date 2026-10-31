/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
 *  All rights reserved.
 *
 * Author: teamgramio (teamgram.io@gmail.com)
 */

package server

import (
	"flag"

	"github.com/azadeveloper18/tgramm/teamgram-server/app/messenger/msg/msg/internal/config"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/messenger/msg/msg/internal/server/grpc"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/messenger/msg/msg/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = flag.String("f", "etc/msg.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	// Modified by Kgramm, 2026-10-31. Original work remains under the Apache License 2.0.
	// UseEnv so Mysql.DSN can read KGRAMM_PG_PASSWORD from Secret kgramm.
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	logx.Infov(c)
	ctx := svc.NewServiceContext(c, nil)
	s.grpcSrv = grpc.New(ctx, c.RpcServerConf)

	go func() {
		go s.grpcSrv.Start()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}
