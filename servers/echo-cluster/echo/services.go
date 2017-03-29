package echo

import (
	"github.com/jennal/goplay/pkg"
	"github.com/jennal/goplay/session"
)

type Services struct {
}

func NewServices() *Services {
	return &Services{}
}

func (self *Services) OnStarted() {
}

func (self *Services) OnStopped() {
}

func (self *Services) OnNewClient(sess *session.Session) {
	sess.Push("echo.push", "Hello from Echo Server")
}

func (self *Services) Echo(sess *session.Session, data string) (string, *pkg.ErrorMessage) {
	return data, nil
}

func (self *Services) Notify(sess *session.Session, data string) *pkg.ErrorMessage {
	sess.Push("echo.push", data)
	return nil
}
