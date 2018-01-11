package main

import (
	"flag"
	"todo/server"
)

var addr = flag.String("addr", ":3000", "Server host")

func main() {
	flag.Parse()

	server.Listen(*addr)
}
