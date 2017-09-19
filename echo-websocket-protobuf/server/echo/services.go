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

package echo

import (
	"github.com/jennal/goplay-demos/echo-websocket-protobuf/proto"
	"github.com/jennal/goplay/log"
	"github.com/jennal/goplay/pkg"
	"github.com/jennal/goplay/session"
	"github.com/jennal/goplay/transfer"
)

type Services struct {
}

func NewServices() *Services {
	return &Services{}
}

func (self *Services) OnStarted() {
}

func (self *Services) OnStopped() {
}

func (self *Services) OnNewClient(sess *session.Session) {
	sess.Once(transfer.EVENT_CLIENT_DISCONNECTED, self, func(client transfer.IClient) {
		log.Logf("Client disconnected: %v\n", client.RemoteAddr())
	})
	log.Tracef("%#v, %#v", sess.Encoding.String(), sess.Encoder)
	sess.Push("echo.push", &proto.String{
		Value: "Hello from Echo Server",
	})
}

func (self *Services) Echo(sess *session.Session, data *proto.String) (*proto.String, *pkg.ErrorMessage) {
	return data, nil
}

func (self *Services) Notify(sess *session.Session, data *proto.String) *pkg.ErrorMessage {
	sess.Push("echo.push", data)
	return nil
}

func (self *Services) Error(sess *session.Session, data *proto.String) (*proto.String, *pkg.ErrorMessage) {
	return nil, pkg.NewErrorMessage(pkg.Status_ERR_WRONG_PARAMS, "STAT_ERR_WRONG_PARAMS")
}
