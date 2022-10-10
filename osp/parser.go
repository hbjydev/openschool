package osp

import (
	"errors"
	"strings"

	"go.h4n.io/openschool/osrn"
)

type OspParser struct{}

func Parse(osp string) (*OspRequest, error) {
	lines := strings.Split(osp, "\n")
	request := OspRequest{}

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
		err = errors.New("not enough request line components")
		return
	}

	action = components[0]
	reqOsrn = osrn.ParseOSRN(components[1])
	version = components[2]
	return
}
