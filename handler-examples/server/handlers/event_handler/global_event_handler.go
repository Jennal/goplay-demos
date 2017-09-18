package event_handler

import (
	"github.com/jennal/goplay/log"
	"github.com/jennal/goplay/session"
)

type GlobalEventHandler struct {
}

func NewGlobalEventHandler() *GlobalEventHandler {
	return &GlobalEventHandler{}
}

func (self *GlobalEventHandler) OnStarted() {
	log.Log("GlobalEventHandler.OnStarted")
}

func (self *GlobalEventHandler) OnStopped() {
	log.Log("GlobalEventHandler.OnStopped")
}

func (self *GlobalEventHandler) OnNewClient(sess *session.Session) {
	log.Logf("GlobalEventHandler.OnNewClient: %v", sess.RemoteAddr())
}
