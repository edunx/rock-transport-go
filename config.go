package transport

import (
	"github.com/edunx/lua"
)

type Proxy struct {
	info   string
	tun    Tunnel
}

type Message interface {
	Byte() []byte
}

type Tunnel interface {
	ToUserData(*lua.LState) *lua.LUserData

	Start() error
	Close()
	Reload()

	Type()  string
	Proxy( string , interface{})

	Push(interface{})
}