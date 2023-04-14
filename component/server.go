//go:build server || full || mini
// +build server full mini

package build

import (
	_ "github.com/jozn/jonz-trojan/proxy/server"
)
