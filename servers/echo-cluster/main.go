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
	"github.com/jennal/goplay-demos/servers/echo-cluster/echo"
	"github.com/jennal/goplay-master/master"
	"github.com/jennal/goplay/cmd"
	"github.com/jennal/goplay/service"
	"github.com/jennal/goplay/transfer/tcp"
)

func main() {
	/* setup service  */
	ser := tcp.NewServer("", 1234)
	serv := service.NewService("echo", ser)

	/* regist handler */
	serv.RegistHanlder(echo.NewServices())

	/* link to master */
	cli := tcp.NewClient()
	mc := master.NewMasterClient(cli)
	sp := master.NewServicePack(master.ST_BACKEND, "echo", 1234)
	mc.Bind(serv, &sp, "", master.PORT)

	cmd.Start(serv)
}
