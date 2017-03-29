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
	err := client.Connect("", 9934)
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
