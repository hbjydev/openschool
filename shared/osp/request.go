package osp

import "go.h4n.io/openschool/shared/osrn"

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
