// bindings.go - Katzenpost client library C binding
// Copyright (C) 2018  David Stainton.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package main for client C binding.
package main

import "C"

import (
	"unsafe"

	"github.com/katzenpost/client"
	"github.com/katzenpost/client/config"
	"github.com/katzenpost/client/session"
)

var myConfig *config.Config
var myClient *client.Client
var mySession *session.Session

//export LoadConfig
func LoadConfig(cfg *C.char) {
	c, err := config.LoadFile(C.GoString(cfg), false)
	if err != nil {
		panic(err)
	}
	myConfig = c
}

//export NewClient
func NewClient() {
	c, err := client.New(myConfig)
	if err != nil {
		panic(err)
	}
	myClient = c
}

//export Start
func Start() {
	s, err := myClient.NewSession()
	if err != nil {
		panic(err)
	}
	mySession = s
}

//export Stop
func Stop() {
	myClient.Shutdown()
}

//export QueryAvailableService
func QueryAvailableService(service *C.char, messagePtr unsafe.Pointer, messageLen C.int) {
	message := C.GoBytes(messagePtr, messageLen)
	serviceDesc, err := mySession.GetService(C.GoString(service))
	if err != nil {
		panic(err)
	}
	msgId, err := mySession.SendUnreliableQuery(serviceDesc.Name, serviceDesc.Provider, message)
	if err != nil {
		panic(err)
	}
	mySession.WaitForReply(msgId)
}

//export QueryService
func QueryService(name, provider *C.char, messagePtr unsafe.Pointer, messageLen C.int) {
	message := C.GoBytes(messagePtr, messageLen)
	msgId, err := mySession.SendUnreliableQuery(C.GoString(name), C.GoString(provider), message)
	if err != nil {
		panic(err)
	}
	mySession.WaitForReply(msgId)
}

func main() {}
