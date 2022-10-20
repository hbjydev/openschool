package main

import (
	"go.h4n.io/openschool/cmd/classes/server"
)

func main() {
	server.NewClassesServerCommand().Execute()
}
