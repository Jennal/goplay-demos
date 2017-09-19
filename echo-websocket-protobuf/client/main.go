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
	"fmt"

	"github.com/jennal/goplay-demos/echo-websocket-protobuf/proto"
	"github.com/jennal/goplay/log"
	"github.com/jennal/goplay/pkg"
	"github.com/jennal/goplay/service"
	"github.com/jennal/goplay/transfer/websocket"
)

func main() {
	cli := websocket.NewClient()
	client := service.NewServiceClient(cli)
	client.SetEncoding(pkg.ENCODING_PROTOBUF)

	client.AddListener("echo.push", func(push *proto.String) {
		log.Log("OnPush: ", push.Value)
	})

	err := client.Connect("", 1234)
	if err != nil {
		log.Error(err)
		return
	}

	for {
		line := ""
		fmt.Scanln(&line)
		log.Log("Send: ", line)

		client.Notify("echo.services.notify", &proto.String{line})

		client.Request("echo.services.echo", &proto.String{
			Value: line,
		}, func(back *proto.String) {
			log.Log("Recv: ", back.Value)
		}, func(err *pkg.ErrorMessage) {
			log.Error(err)
		})
	}
}
