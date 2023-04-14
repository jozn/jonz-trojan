package transport

import (
	"net"

	"github.com/jozn/jonz-trojan/tunnel"
)

type Conn struct {
	net.Conn
}

func (c *Conn) Metadata() *tunnel.Metadata {
	return nil
}
