package osp

import (
	"strings"

	"go.h4n.io/openschool/osrn"
)

type OspParser struct{}

func Parse(osp string) (*Request, error) {
	lines := strings.Split(osp, "\n")
	request := Request{}

	if len(lines) == 0 {
		return nil, ErrorBadVersion
	}

	rawAction, osrn, ver, err := parseRequestLine(lines[0])
	if err != nil {
		return nil, err
	}

	action, ok := GetAction(rawAction)
	if !ok {
		return nil, ErrorBadAction
	}

	request.Action = action
	request.Osrn = osrn
	request.Version = ver

	return &request, nil
}

func parseRequestLine(line string) (action string, reqOsrn osrn.OSRN, version string, err error) {
	components := strings.Split(line, " ")

	if len(components) != 3 {
		err = ErrorNotEnoughRequestLineComponents
		return
	}

	action = components[0]
	reqOsrn = osrn.ParseOSRN(components[1])
	version = components[2]
	return
}
