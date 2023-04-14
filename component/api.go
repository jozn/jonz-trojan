//go:build api || full
// +build api full

package build

import (
	_ "github.com/jozn/jonz-trojan/api/control"
	_ "github.com/jozn/jonz-trojan/api/service"
)
