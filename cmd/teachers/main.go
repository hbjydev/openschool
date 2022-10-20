package main

import (
	"go.h4n.io/openschool/cli"
	"go.h4n.io/openschool/cmd/teachers/server"
)

func main() {
	cli.Run(server.NewTeachersServerCommand())
}
