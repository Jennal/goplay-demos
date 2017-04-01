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
	"bufio"
	"fmt"
	"os"

	"time"

	"strings"

	"github.com/jennal/goplay-demos/chat/server/chat"
	"github.com/jennal/goplay/aop"
	"github.com/jennal/goplay/log"
	"github.com/jennal/goplay/pkg"
	"github.com/jennal/goplay/service"
	"github.com/jennal/goplay/transfer/tcp"
)

const (
	RoomName = "Room-1"
)

func main() {
	cli := tcp.NewClient()
	client := service.NewServiceClient(cli)

	client.AddListener("chat.push", func(push string) {
		log.Log("OnPush: ", push)
	})

	client.AddListener("channel."+strings.ToLower(RoomName), func(push string) {
		log.Log("OnChannel: ", push)
	})

	err := client.Connect("", 2234)
	if err != nil {
		log.Error(err)
		return
	}

	time.Sleep(100 * time.Millisecond)

	var name string
	var rooms []string
	fmt.Printf("Enter Name: ")
	fmt.Scanln(&name)

	//Get Rooms
	aop.Parallel(func(done chan bool) {
		client.Request("chat.services.rooms", 0, func(rs []string) {
			rooms = rs
			fmt.Println("Rooms on server:")
			for _, r := range rooms {
				fmt.Println("\t", r)
			}
			done <- true
		}, func(err *pkg.ErrorMessage) {
			log.Error(err)
			done <- true
		})
	})

	isCreateRoom := true
	for _, r := range rooms {
		if r == RoomName {
			isCreateRoom = false
			break
		}
	}

	if isCreateRoom {
		aop.Parallel(func(done chan bool) {
			client.Request("chat.services.create", RoomName, func(st pkg.Status) {
				fmt.Printf("Room \"%v\" Created!\n", RoomName)
				done <- true
			}, func(err *pkg.ErrorMessage) {
				log.Error(err)
				done <- true
			})
		})
	} else {
		aop.Parallel(func(done chan bool) {
			client.Request("chat.services.join", RoomName, func(st pkg.Status) {
				fmt.Printf("Join Room \"%v\" Successed!\n", RoomName)
				done <- true
			}, func(err *pkg.ErrorMessage) {
				log.Error(err)
				done <- true
			})
		})
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		client.Request("chat.services.say", chat.ChatData{
			UserName: name,
			RoomName: RoomName,
			Message:  scanner.Text(),
		}, func(back string) {
			log.Log(back)
		}, func(err *pkg.ErrorMessage) {
			log.Error(err)
		})
	}
}
