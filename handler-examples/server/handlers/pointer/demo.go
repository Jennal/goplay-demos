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

package pointer

import (
	"github.com/jennal/goplay/pkg"
	"github.com/jennal/goplay/session"
)

const (
	PUSH_KEY = "pointer.demo.push"
)

type Demo struct {
}

func NewDemo() *Demo {
	return &Demo{}
}

func (self *Demo) OnStarted() {
}

func (self *Demo) OnStopped() {
}

func (self *Demo) OnNewClient(sess *session.Session) {

}

func (self *Demo) RequestString(sess *session.Session, data *string) (*string, *pkg.ErrorMessage) {
	return data, nil
}

func (self *Demo) NotifyString(sess *session.Session, data *string) *pkg.ErrorMessage {
	sess.Push(PUSH_KEY, data)
	return nil
}

func (self *Demo) RequestError(sess *session.Session) (string, *pkg.ErrorMessage) {
	return "", pkg.NewErrorMessage(pkg.STAT_ERR_WRONG_PARAMS, "STAT_ERR_WRONG_PARAMS")
}
