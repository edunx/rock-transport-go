package transport

import (
	"github.com/edunx/lua"
)

const (
	TPPROXYMT string = "ROCK_TRANSPORT_PROXY_GO_MT"
)

func (p *Proxy) ToUserData(L *lua.LState) *lua.LUserData {
	return L.NewUserDataByInterface(p, TPPROXYMT)
}

func luaInjectTransportProxyApi(L *lua.LState , parent *lua.LTable) {
	mt := L.NewTypeMetatable( TPPROXYMT )
	L.SetField(mt ,"__index" , L.NewFunction(ProxyGet))
	L.SetField(mt , "__newindex" , L.NewFunction(ProxySet))

	L.SetField(parent , "proxy" , L.NewFunction(CreateTransportProxyUserData))
}

func CreateTransportProxyUserData(L *lua.LState) int {
	opt := L.CheckTable(1)

	v := &Proxy{
		info: opt.CheckString("info" , "unkown"),
		tun: CheckTunnelUserDataByTable(L , opt , "tunnel"),
	}

	ud := L.NewUserDataByInterface( v , TPPROXYMT )

	L.Push(ud)
	return 1
}

func ProxyGet(L *lua.LState) int {
	return 0
}

func ProxySet(L *lua.LState) int {
	return 0
}

func (p *Proxy) Push( data interface{} ) {
}