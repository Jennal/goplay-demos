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

	"github.com/jennal/goplay-demos/handler-examples/server/handlers/pointer"
	"github.com/jennal/goplay-demos/handler-examples/server/handlers/push"
	"github.com/jennal/goplay/log"
	"github.com/jennal/goplay/pkg"
	"github.com/jennal/goplay/service"
	"github.com/jennal/goplay/transfer/tcp"
)

func main() {
	cli := tcp.NewClient()
	client := service.NewServiceClient(cli)

	client.AddListener(push.PUSH_KEY, func(push string) {
		log.Log("OnPush[push.echo]: ", push)
	})

	client.AddListener(pointer.PUSH_KEY, func(push string) {
		log.Logf("OnPush[%v]: %v", pointer.PUSH_KEY, push)
	})

	client.AddListener(pointer.PUSH_KEY, func(push *string) {
		log.Logf("OnPush[%v]: (%p)%v", pointer.PUSH_KEY, push, *push)
	})

	err := client.Connect("", 2468)
	if err != nil {
		log.Error(err)
		return
	}

	for {
		fmt.Print("Input: ")
		line := ""
		fmt.Scanln(&line)
		log.Log("Send: ", line)

		testPointer(client, line)
	}
}

func testPointer(client *service.ServiceClient, line string) {
	client.Notify("pointer.demo.notifystring", line)
	client.Request("pointer.demo.requeststring", line, func(back *string) {
		log.Log("Recv[pointer.demo.requeststring]: ", back, "\t", *back)
	}, func(err *pkg.ErrorMessage) {
		log.Error(err)
	})
	client.Request("pointer.demo.requesterror", nil, func(back string) {
		log.Log("Recv[pointer.demo.requesterror]: ", back)
	}, func(err *pkg.ErrorMessage) {
		log.Errorf("Recv[pointer.demo.requesterror]: %v", err)
	})
}
