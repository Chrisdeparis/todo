package main

import "todo/server"

// export GOPATH=/home/dev/go
// go get github.com/satori/go.uuid
// go get github.com/gorilla/mux
func main() {
	server.Listen(":7070")
}
