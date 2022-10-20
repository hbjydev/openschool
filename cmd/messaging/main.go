package main

import (
	"go.h4n.io/openschool/cmd/messaging/server"
)

func main() {
	server.NewMessagingServerCommand().Execute()
}
