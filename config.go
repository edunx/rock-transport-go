package transport

import (
	"github.com/edunx/lua"
)

type Proxy struct {
	info   string
	T      int
	tun    Tunnel
}

type Tunnel interface {

	ToUserData(*lua.LState) *lua.LUserData

	Start() error
	Close()
	Reload()

	Push(interface{})
}