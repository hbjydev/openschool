package main

import (
	"go.h4n.io/openschool/cli"
	"go.h4n.io/openschool/cmd/students/server"
)

func main() {
	cli.Run(server.NewStudentsServerCommand())
}
