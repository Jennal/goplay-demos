// Copyright (C) 2017 Jennal(jennalcn@gmail.com). All rights reserved.
//
// Licensed under the MIT License (the "License"); you may not use this file except
// in compliance with the License. You may obtain a copy of the License at
//
// http://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package chat

import (
	"fmt"

	"strings"

	"github.com/jennal/goplay/channel"
	"github.com/jennal/goplay/log"
	"github.com/jennal/goplay/pkg"
	"github.com/jennal/goplay/session"
	"github.com/jennal/goplay/transfer"
)

type Services struct {
	cm *channel.ChannelManager
}

func NewServices() *Services {
	return &Services{
		cm: channel.GetChannelManager(),
	}
}

func (self *Services) OnStarted() {
}

func (self *Services) OnStopped() {
}

func (self *Services) OnNewClient(sess *session.Session) {
	sess.Once(transfer.EVENT_CLIENT_DISCONNECTED, self, func(client transfer.IClient) {
		log.Logf("Client disconnected: %v\n", client.RemoteAddr())
	})
	sess.Push("chat.push", "Hello from Chat Server")
	sess.Set("TEST", "OK")
	log.Logf("====> OnNewClient: %p", sess)
}

func (self *Services) Rooms(sess *session.Session, ignore int) ([]string, *pkg.ErrorMessage) {
	log.Logf("====> Rooms: %p", sess)
	log.Log(sess.String("TEST"))
	return self.cm.ChannelNames(), nil
}

func (self *Services) Create(sess *session.Session, roomName string) (pkg.Status, *pkg.ErrorMessage) {
	roomName = strings.TrimSpace(roomName)
	if roomName == "" {
		return pkg.STAT_ERR, pkg.NewErrorMessage(pkg.STAT_ERR, "room name should not be empty")
	}

	ch := self.cm.Create(roomName)
	ch.Add(sess)
	log.Log("===> ", ch.Count(), "\t", sess)

	return pkg.STAT_OK, nil
}

func (self *Services) Join(sess *session.Session, roomName string) (pkg.Status, *pkg.ErrorMessage) {
	roomName = strings.TrimSpace(roomName)
	if roomName == "" {
		return pkg.STAT_ERR, pkg.NewErrorMessage(pkg.STAT_ERR, "room name should not be empty")
	}

	ch := self.cm.Get(roomName)
	if ch == nil {
		return pkg.STAT_ERR, pkg.NewErrorMessage(pkg.STAT_ERR, "room not exists")
	}

	ch.Add(sess)
	log.Log("===> ", ch.Count(), "\t", sess)
	return pkg.STAT_OK, nil
}

func (self *Services) Say(sess *session.Session, data ChatData) *pkg.ErrorMessage {
	ch := self.cm.Get(data.RoomName)
	if ch == nil {
		return pkg.NewErrorMessage(pkg.STAT_ERR, "room not exists")
	}

	if !ch.Exists(sess) {
		return pkg.NewErrorMessage(pkg.STAT_ERR, "user not exists in room")
	}

	ch.Broadcast(fmt.Sprintf("%v: %v", data.UserName, data.Message))
	return nil
}
