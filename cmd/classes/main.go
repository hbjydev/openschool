package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/lucsky/cuid"
	"go.h4n.io/openschool/class/models"
	"go.h4n.io/openschool/class/repos/class"
	"go.h4n.io/openschool/osp"
)

func main() {
	// cert, err := tls.LoadX509KeyPair("pki/classes/fullchain.pem", "pki/classes/cert.key")
	// if err != nil {
	// panic(err)
	// }
	// cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	repo := class.InMemoryClassRepository{
		Items: []models.Class{
			{
				Id:          cuid.New(),
				Name:        `class-1a`,
				DisplayName: `Class 1a: Mathematics 101`,
				Description: `A mathematics class`,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Id:          cuid.New(),
				Name:        `class-1b`,
				DisplayName: `Class 1b: Chemistry 101`,
				Description: `A chemistry class`,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				Id:          cuid.New(),
				Name:        `class-1c`,
				DisplayName: `Class 1c: Physics 101`,
				Description: `A physics class`,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
	}

	classResource := osp.Resource{
		LIST: func(request *osp.OspRequest) (osp.Response, error) {
			items, err := repo.GetAll(10, 1)
			if err != nil {
				return osp.Response{}, err
			}

			body, err := json.Marshal(items)
			if err != nil {
				return osp.Response{}, err
			}

			return osp.Response{
				Status: osp.OspStatusSuccess,
				Headers: map[string]string{
					"content-type": "text/plain",
				},
				Body: string(body),
			}, nil
		},
	}

	server := &osp.Service{
		Addr: `0.0.0.0:8001`,
		Name: `classes`,
		Resources: map[string]osp.Resource{
			"class": classResource,
		},
		// Tls:  cfg,
	}

	log.Fatal(server.Run())
}
