package main

import (
	"go.h4n.io/openschool/cmd/terms/server"
)

func main() {
	server.NewTermsServerCommand().Execute()
}
