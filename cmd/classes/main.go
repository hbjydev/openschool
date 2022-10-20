package main

import (
	"go.h4n.io/openschool/cli"
	"go.h4n.io/openschool/cmd/classes/server"
)

func main() {
	cli.Run(server.NewClassesServerCommand())
}
