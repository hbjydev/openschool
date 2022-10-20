package osp

import (
	"fmt"

	"go.h4n.io/openschool/osrn"
)

type Action string

const (
	ActionGet    Action = "GET"
	ActionList   Action = "LIST"
	ActionUpdate Action = "UPDATE"
	ActionCreate Action = "CREATE"
	ActionDelete Action = "DELETE"
)

var (
	ospActionMap = map[string]Action{
		"GET":    ActionGet,
		"LIST":   ActionList,
		"UPDATE": ActionUpdate,
		"CREATE": ActionCreate,
		"DELETE": ActionDelete,
	}
)

func GetAction(action string) (Action, bool) {
	a, ok := ospActionMap[action]
	return a, ok
}

type Request struct {
	Action  Action
	Osrn    osrn.OSRN
	Version string
	Headers map[string]string
	Body    string
}

func (r *Request) LogMaps() []interface{} {
	return []interface{}{
		"request.action", r.Action,
		"request.osrn", r.Osrn,
		"request.version", r.Version,
	}
}

func (r *Request) HeadersString() string {
	lines := ``

	for k, v := range r.Headers {
		line := fmt.Sprintf("%v: %v", k, v)
		lines = fmt.Sprintf("%v%v\r\n", lines, line)
	}

	return lines
}

func (r *Request) String() string {
	request := fmt.Sprintf("%v %v %v", r.Action, r.Osrn, r.Version)

	if len(r.Headers) > 0 {
		request = fmt.Sprintf("%v\r\n%v", request, r.HeadersString())
	}

	if len(r.Body) > 0 {
		request = fmt.Sprintf("%v\r\n%v", request, r.Body)
	}

	return fmt.Sprintf("%v", request)
}

func (r *Request) Bytes() []byte {
	return []byte(r.String())
}
