package main

import (
	"go.h4n.io/openschool/cmd/parents/server"
)

func main() {
	server.NewParentsServerCommand().Execute()
}
