package main

import (
	"go.h4n.io/openschool/cmd/teachers/server"
)

func main() {
	server.NewTeachersServerCommand().Execute()
}
