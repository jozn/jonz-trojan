package main

import (
	"flag"

	_ "github.com/jozn/jonz-trojan/component"
	"github.com/jozn/jonz-trojan/log"
	"github.com/jozn/jonz-trojan/option"
)

func main() {
	flag.Parse()
	for {
		h, err := option.PopOptionHandler()
		if err != nil {
			log.Fatal("invalid options")
		}
		err = h.Handle()
		if err == nil {
			break
		}
	}
}
