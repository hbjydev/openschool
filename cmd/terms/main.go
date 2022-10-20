package main

import (
	"go.h4n.io/openschool/cli"
	"go.h4n.io/openschool/cmd/terms/server"
)

func main() {
	cli.Run(server.NewTermsServerCommand())
}
