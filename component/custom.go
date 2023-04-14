//go:build custom || full
// +build custom full

package build

import (
	_ "github.com/jozn/jonz-trojan/proxy/custom"
	_ "github.com/jozn/jonz-trojan/tunnel/adapter"
	_ "github.com/jozn/jonz-trojan/tunnel/dokodemo"
	_ "github.com/jozn/jonz-trojan/tunnel/freedom"
	_ "github.com/jozn/jonz-trojan/tunnel/http"
	_ "github.com/jozn/jonz-trojan/tunnel/mux"
	_ "github.com/jozn/jonz-trojan/tunnel/router"
	_ "github.com/jozn/jonz-trojan/tunnel/shadowsocks"
	_ "github.com/jozn/jonz-trojan/tunnel/simplesocks"
	_ "github.com/jozn/jonz-trojan/tunnel/socks"
	_ "github.com/jozn/jonz-trojan/tunnel/tls"
	_ "github.com/jozn/jonz-trojan/tunnel/tproxy"
	_ "github.com/jozn/jonz-trojan/tunnel/transport"
	_ "github.com/jozn/jonz-trojan/tunnel/trojan"
	_ "github.com/jozn/jonz-trojan/tunnel/websocket"
)
