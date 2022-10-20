package server

import (
	"encoding/json"

	"go.h4n.io/openschool/class/repos/class"
	"go.h4n.io/openschool/osp"
	"go.h4n.io/openschool/osp/osputil"
)

func NewClassResource(repo class.ClassRepository) osp.Resource {
	return osp.Resource{
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
	}
}
