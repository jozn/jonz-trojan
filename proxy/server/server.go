package server

import (
	"context"

	"github.com/jozn/jonz-trojan/config"
	"github.com/jozn/jonz-trojan/proxy"
	"github.com/jozn/jonz-trojan/proxy/client"
	"github.com/jozn/jonz-trojan/tunnel/freedom"
	"github.com/jozn/jonz-trojan/tunnel/mux"
	"github.com/jozn/jonz-trojan/tunnel/router"
	"github.com/jozn/jonz-trojan/tunnel/shadowsocks"
	"github.com/jozn/jonz-trojan/tunnel/simplesocks"
	"github.com/jozn/jonz-trojan/tunnel/tls"
	"github.com/jozn/jonz-trojan/tunnel/transport"
	"github.com/jozn/jonz-trojan/tunnel/trojan"
	"github.com/jozn/jonz-trojan/tunnel/websocket"
)

const Name = "SERVER"

func init() {
	proxy.RegisterProxyCreator(Name, func(ctx context.Context) (*proxy.Proxy, error) {
		cfg := config.FromContext(ctx, Name).(*client.Config)
		ctx, cancel := context.WithCancel(ctx)
		transportServer, err := transport.NewServer(ctx, nil)
		if err != nil {
			cancel()
			return nil, err
		}
		clientStack := []string{freedom.Name}
		if cfg.Router.Enabled {
			clientStack = []string{freedom.Name, router.Name}
		}

		root := &proxy.Node{
			Name:       transport.Name,
			Next:       make(map[string]*proxy.Node),
			IsEndpoint: false,
			Context:    ctx,
			Server:     transportServer,
		}

		if !cfg.TransportPlugin.Enabled {
			root = root.BuildNext(tls.Name)
		}

		trojanSubTree := root
		if cfg.Shadowsocks.Enabled {
			trojanSubTree = trojanSubTree.BuildNext(shadowsocks.Name)
		}
		trojanSubTree.BuildNext(trojan.Name).BuildNext(mux.Name).BuildNext(simplesocks.Name).IsEndpoint = true
		trojanSubTree.BuildNext(trojan.Name).IsEndpoint = true

		wsSubTree := root.BuildNext(websocket.Name)
		if cfg.Shadowsocks.Enabled {
			wsSubTree = wsSubTree.BuildNext(shadowsocks.Name)
		}
		wsSubTree.BuildNext(trojan.Name).BuildNext(mux.Name).BuildNext(simplesocks.Name).IsEndpoint = true
		wsSubTree.BuildNext(trojan.Name).IsEndpoint = true

		serverList := proxy.FindAllEndpoints(root)
		clientList, err := proxy.CreateClientStack(ctx, clientStack)
		if err != nil {
			cancel()
			return nil, err
		}
		return proxy.NewProxy(ctx, cancel, serverList, clientList), nil
	})
}
