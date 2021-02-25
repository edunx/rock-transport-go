package transport

import (
	pub "github.com/edunx/rock-public-go"
)

type Proxy struct {
	info   string
	T      int
	tun    pub.Transport
}