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

	kafka "github.com/teamgram/marmota/pkg/mq"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/config"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/server/mq"
	"github.com/teamgram/teamgram-server/app/messenger/sync/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/sync.yaml", "the config file")

type Server struct {
	// grpcSrv *zrpc.RpcServer
	mq *kafka.ConsumerGroup
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

	if err := logx.SetUp(c.Log); err != nil {
		return err
	}

	ctx := svc.NewServiceContext(c)
	// s.grpcSrv = grpc.New(ctx, c.RpcServerConf)
	s.mq = mq.New(ctx, c.SyncConsumer)

	// go s.grpcSrv.Start()
	go s.mq.Start()

	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	// s.grpcSrv.Stop()
	s.mq.Stop()
}
