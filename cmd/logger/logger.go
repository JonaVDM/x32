package main

import (
	"github.com/jonavdm/x32/pkg/x32"
)

func main() {
	// addr := "192.168.0.101:10023"
	addr := "10.0.0.21:10023"

	c := x32.NewClient(addr)
	err := c.Connect()
	defer c.Close()

	if err != nil {
		panic(err)
	}

	c.Logger()
}
