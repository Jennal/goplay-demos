package main

import (
	"github.com/jennal/goplay-demos/servers/echo/echo"
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

	cmd.Start(serv)
}
