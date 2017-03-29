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

	"github.com/jennal/goplay/log"
	"github.com/jennal/goplay/pkg"
	"github.com/jennal/goplay/service"
	"github.com/jennal/goplay/transfer/tcp"
)

func main() {
	cli := tcp.NewClient()
	client := service.NewServiceClient(cli)
	err := client.Connect("", 1234)
	if err != nil {
		log.Error(err)
		return
	}

	client.AddListener("echo.push", func(push string) {
		log.Log("OnPush: ", push)
	})

	for {
		line := ""
		fmt.Scanln(&line)
		log.Log("Send: ", line)

		client.Notify("echo.services.notify", line)

		client.Request("echo.services.echo", line, func(back string) {
			log.Log("Recv: ", back)
		}, func(err *pkg.ErrorMessage) {
			log.Error(err)
		})
	}
}
