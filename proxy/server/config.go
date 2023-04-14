package server

import (
	"github.com/jozn/jonz-trojan/config"
	"github.com/jozn/jonz-trojan/proxy/client"
)

func init() {
	config.RegisterConfigCreator(Name, func() any {
		return new(client.Config)
	})
}
