package main

import (
	"log"

	"go.h4n.io/openschool/osp"
)

func main() {
	server := &osp.Service{
		Addr: `0.0.0.0:8004`,
		Name: `students`,
	}

	// classRepo := class.InMemoryClassRepository{
	// Items: []models.Class{},
	// }

	log.Fatal(server.Run())
}
