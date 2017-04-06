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
	"flag"

	"github.com/jennal/goplay-demos/chat/server/chat"
	"github.com/jennal/goplay-master/master"
	"github.com/jennal/goplay/cmd"
	"github.com/jennal/goplay/transfer/tcp"
)

var (
	port *int
)

func init() {
	port = flag.Int("p", 2234, "port of the server")
}

func main() {
	flag.Parse()

	/* setup service  */
	ser := tcp.NewServer("", *port)
	serv := master.NewBackendService("chat", ser)

	/* regist handler */
	serv.RegistHanlder(chat.NewServices())

	/* link to master */
	serv.ConnectMaster("", master.PORT)

	cmd.Start(serv)
}
