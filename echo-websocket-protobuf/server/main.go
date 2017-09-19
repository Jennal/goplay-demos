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
	"github.com/jennal/goplay-demos/echo-websocket-protobuf/server/echo"
	"github.com/jennal/goplay/cmd"
	"github.com/jennal/goplay/pkg"
	"github.com/jennal/goplay/service"
	"github.com/jennal/goplay/transfer/websocket"
)

func main() {
	/* setup service  */
	ser := websocket.NewServer("", 1234)
	serv := service.NewService("echo", ser)
	serv.SetEncoding(pkg.ENCODING_PROTOBUF)

	/* regist handler */
	serv.RegistHanlder(echo.NewServices())

	cmd.Start(serv)
}
