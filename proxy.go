package transport

import "fmt"

func (p *Proxy) Type() string {
	return fmt.Sprintf("proxy:%s" , p.info)
}
func (p *Proxy) Push( data interface{} ) {
	p.tun.Proxy( p.info , data ) //打上 PROXY 协议标签
}

func (p *Proxy) Proxy(info string , v interface{}) {
	p.tun.Proxy(info , v)
}
