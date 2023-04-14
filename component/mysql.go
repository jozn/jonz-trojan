//go:build mysql || full || mini
// +build mysql full mini

package build

import (
	_ "github.com/jozn/jonz-trojan/statistic/mysql"
)
