package transport

import (
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
)

func CheckTunnel(ud *lua.LUserData) Tunnel {
	switch tun := ud.Value.(type) {
	case Tunnel:
		return tun
	default:
		return nil
	}
}

func CheckTunnelUserData(L *lua.LState , idx int) Tunnel {
	ud := pub.CheckUserData(L , idx)

	tun := CheckTunnel( ud.ToUserData(L) )
	if tun == nil {
		L.RaiseError("invalid type , #%d must be transport.Tunnel " , idx)
		return nil
	}

	return tun
}

func CheckTunnelUserDataByTable(L *lua.LState , tab *lua.LTable , key string) Tunnel {

	ud , ok := tab.RawGetString(key).(*lua.LUserData)
	if !ok {
		L.RaiseError("invalid type , tab %s not userdata" , key)
		return nil
	}

	tun , ok := ud.Value.(Tunnel)
	if !ok {
		L.RaiseError("invalid type , opt %s not transport tunnel" , key)
		return nil

	}

	return tun
}
