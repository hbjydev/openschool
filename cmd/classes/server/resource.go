package server

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"go.h4n.io/openschool/bus"
	"go.h4n.io/openschool/class/repos/class"
	"go.h4n.io/openschool/osp"
	"go.h4n.io/openschool/osp/osputil"
)

func NewClassResource(repo class.ClassRepository) osp.Resource {
	return osp.Resource{
		GET: func(request *osp.Request) (osp.Response, error) {
			id := request.Osrn.Id

			ch, err := bus.BusInst.Channel()
			if err != nil {
				return osp.Response{}, err
			}

      err = ch.ExchangeDeclare("class.read", "topic", true, false, false, false, nil)
			if err != nil {
				return osp.Response{}, err
			}

      err = ch.Publish("classes", "read", true, true, amqp.Publishing{
        ContentType: "application/json",
        Body: []byte(`{"hello":"world"}`),
      })
      if err != nil {
        return osp.Response{}, err
      }

			class, err := repo.Get(id)
			if err != nil {
				return osp.Response{}, err
			}

			if class == nil {
				return osp.Response{
					Status: osp.OspStatusNotFound,
				}, nil
			}

			classJson, err := json.Marshal(class)
			if err != nil {
				return osp.Response{}, err
			}

			return osp.Response{
				Status: osp.OspStatusSuccess,
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: string(classJson),
			}, nil
		},

		LIST: func(request *osp.Request) (osp.Response, error) {
			page, perPage, err := osputil.PaginationFromRequest(request)
			if err != nil {
				return osp.Response{}, err
			}

			items, err := repo.GetAll(perPage, page)
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

		CREATE: func(request *osp.Request) (osp.Response, error) {
			return osp.Response{
				Status: osp.OspStatusCreated,
				Headers: map[string]string{
					"content-type": "text/plain",
				},
				Body: "Did the thing",
			}, nil
		},
	}
}
