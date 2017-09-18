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

package main

import (
	"github.com/jennal/goplay-demos/handler-examples/server/handlers/event_handler"
	"github.com/jennal/goplay-demos/handler-examples/server/handlers/pointer"
	"github.com/jennal/goplay-demos/handler-examples/server/handlers/push"

	"github.com/jennal/goplay/cmd"
	"github.com/jennal/goplay/handler"
	"github.com/jennal/goplay/service"
	"github.com/jennal/goplay/transfer/tcp"
)

func main() {
	/* setup service  */
	ser := tcp.NewServer("", 2468)
	serv := service.NewService("", ser)

	/* regist handler */
	serv.RegistHanlderGroup(map[string][]handler.IHandler{
		"-": []handler.IHandler{
			event_handler.NewGlobalEventHandler(),
		},
		"push": []handler.IHandler{
			push.NewDemo(),
		},
		"sysvalue": []handler.IHandler{},
		"struct":   []handler.IHandler{},
		"pointer": []handler.IHandler{
			pointer.NewDemo(),
		},
		"rawvalue": []handler.IHandler{},
	})

	cmd.Start(serv)
}
