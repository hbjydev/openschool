package main

import (
	"log"

	"go.h4n.io/openschool/osp"
)

func main() {
	// cert, err := tls.LoadX509KeyPair("pki/classes/fullchain.pem", "pki/classes/cert.key")
	// if err != nil {
	// panic(err)
	// }
	// cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	classResource := osp.Resource{
		LIST: func(request *osp.OspRequest) (osp.Response, error) {
			return osp.Response{
				Status: osp.OspStatusSuccess,
				Headers: map[string]string{
					"content-type": "text/plain",
				},
				Body: `list response`,
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
