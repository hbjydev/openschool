package osp

import (
	"fmt"

	"go.h4n.io/openschool/osrn"
)

type OspAction string

const (
	OspActionGet    OspAction = "GET"
	OspActionList   OspAction = "LIST"
	OspActionUpdate OspAction = "UPDATE"
	OspActionCreate OspAction = "CREATE"
	OspActionDelete OspAction = "DELETE"
)

var (
	ospActionMap = map[string]OspAction{
		"GET":    OspActionGet,
		"LIST":   OspActionList,
		"UPDATE": OspActionUpdate,
		"CREATE": OspActionCreate,
		"DELETE": OspActionDelete,
	}
)

func GetAction(action string) (OspAction, bool) {
	a, ok := ospActionMap[action]
	return a, ok
}

type OspRequest struct {
	Action  OspAction
	Osrn    osrn.OSRN
	Version string
}

func (r *OspRequest) LogMaps() []interface{} {
	return []interface{}{
		"request.action", r.Action,
		"request.osrn", r.Osrn,
		"request.version", r.Version,
	}
}

func (r *OspRequest) String() string {
	return fmt.Sprintf("%v %v %v\n\n", r.Action, r.Osrn, r.Version)
}

func (r *OspRequest) Bytes() []byte {
	return []byte(r.String())
}
