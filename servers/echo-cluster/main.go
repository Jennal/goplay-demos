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
