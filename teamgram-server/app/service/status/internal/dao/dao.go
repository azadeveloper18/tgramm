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
	"github.com/azadeveloper18/tgramm/marmota/pkg/stores/kv"
	"github.com/azadeveloper18/tgramm/teamgram-server/app/service/status/internal/config"
)

type Dao struct {
	KV kv.ExtStore
}

func New(c config.Config) *Dao {
	d := &Dao{
		KV: kv.NewStore(c.Status),
	}

	return d
}
