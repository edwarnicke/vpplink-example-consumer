// Code generated by vpplink DO NOT EDIT.
// Copyright (C) 2021 Cisco Systems Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpplink

import (
	"sync"
	"time"

	govpp "go.fd.io/govpp"
	vppapi "go.fd.io/govpp/api"
	vppcore "go.fd.io/govpp/core"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	DefaultReplyTimeout = 1 * time.Second
	InvalidSwIfIndex    = ^uint32(1)
)

// Vpp is the base struct that exposes all the methods defined
// by the various wrappers.
// Depending on the available APIs, this struct will implement
// the various interfaces defined in github.com/vpplink/api
type Vpp struct {
	lock   sync.Mutex
	conn   *vppcore.Connection
	ch     vppapi.Channel
	socket string
	log    *logrus.Entry
}

func (v *Vpp) GetLog() *logrus.Entry {
	return v.log
}

func (v *Vpp) GetChannel() vppapi.Channel {
	return v.ch
}

func (v *Vpp) Lock() {
	v.lock.Lock()
}

func (v *Vpp) Unlock() {
	v.lock.Unlock()
}

func (v *Vpp) MakeNewChannel() (vppapi.Channel, error) {
	return v.conn.NewAPIChannel()
}

func NewVpp(socket string, logger *logrus.Entry) (*Vpp, error) {
	conn, err := govpp.Connect(socket)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot connect to VPP on socket %s", socket)
	}

	ch, err := conn.NewAPIChannel()
	if err != nil {
		return nil, errors.Wrap(err, "channel creation failed")
	}

	return &Vpp{
		conn:   conn,
		ch:     ch,
		socket: socket,
		log:    logger,
	}, nil
}

func (v *Vpp) Reconnect() (err error) {
	v.conn, err = govpp.Connect(v.socket)
	if err != nil {
		return errors.Wrapf(err, "cannot re-connect to VPP on socket %s", v.socket)
	}
	v.ch, err = v.conn.NewAPIChannel()
	if err != nil {
		return errors.Wrap(err, "channel re-creation failed")
	}
	return nil
}

func (v *Vpp) Close() error {
	if v == nil {
		return nil
	}
	if v.ch != nil {
		v.ch.Close()
	}
	if v.conn != nil {
		v.conn.Disconnect()
	}
	return nil
}

func (v *Vpp) SendRequestAwaitReply(request vppapi.Message, response vppapi.Message) error {
	v.Lock()
	defer v.Unlock()

	err := v.GetChannel().SendRequest(request).ReceiveReply(response)
	if err != nil {
		return errors.Wrapf(err, "API internal error, msg=%s", request.GetMessageName())
	}

	return nil
}
