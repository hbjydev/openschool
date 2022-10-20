package osputil

import (
	"strconv"

	"go.h4n.io/openschool/osp"
)

func PaginationFromRequest(req *osp.Request) (page int, perPage int, err error) {
	perPageStr, ok := req.Headers["per-page"]
	if !ok {
		perPageStr = "10"
	}

	pageStr, ok := req.Headers["page"]
	if !ok {
		pageStr = "1"
	}

	perPage, err = strconv.Atoi(perPageStr)
	if err != nil {
		return
	}
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		return
	}

	return
}
