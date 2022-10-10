package main

import (
	"log"

	"go.h4n.io/openschool/osp"
)

func main() {
	server := &osp.Service{
		Addr: `0.0.0.0:8002`,
		Name: `messaging`,
	}

	// classRepo := class.InMemoryClassRepository{
	// Items: []models.Class{},
	// }

	log.Fatal(server.Run())
}
