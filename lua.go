package transport

import "github.com/edunx/lua"

func LuaInjectApi(L *lua.LState , parent *lua.LTable) {
	transportTab := L.CreateTable(0 , 1)
	luaInjectTransportProxyApi(L , transportTab)
	L.SetField(parent , "transport" , transportTab)
}
