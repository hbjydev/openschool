package main

import (
	"go.h4n.io/openschool/cmd/students/server"
)

func main() {
	server.NewStudentsServerCommand().Execute()
}
