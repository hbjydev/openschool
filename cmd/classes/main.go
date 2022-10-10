package main

import (
	"log"

	"go.h4n.io/openschool/osp"
)

func main() {
	server := &osp.Service{
		Addr: `0.0.0.0:8001`,
		Name: `classes`,
	}

	log.Fatal(server.Run())
}
