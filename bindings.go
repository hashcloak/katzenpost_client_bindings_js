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
)

var myConfig *config.Config
var myClient *client.Client
var mySession *client.Session

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
func QueryAvailableService(service *C.char, messagePtr unsafe.Pointer, messageLen C.int) unsafe.Pointer {
	message := C.GoBytes(messagePtr, messageLen)
	serviceDesc, err := mySession.GetService(C.GoString(service))
	if err != nil {
		panic(err)
	}
	reply, err := mySession.SendUnreliableMessage(serviceDesc.Name, serviceDesc.Provider, message)
	if err != nil {
		panic(err)
	}
	return C.CBytes(reply)
}

//export SendUnreliableMessage
func SendUnreliableMessage(name, provider *C.char, messagePtr unsafe.Pointer, messageLen C.int) unsafe.Pointer {
	message := C.GoBytes(messagePtr, messageLen)
	reply, err := mySession.SendUnreliableMessage(C.GoString(name), C.GoString(provider), message)
	if err != nil {
		panic(err)
	}
	return C.CBytes(reply)
}

func main() {}
