package osp_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"go.h4n.io/openschool/shared/osp"
	"go.h4n.io/openschool/shared/osrn"
)

func TestParseOnlyRequestLine(t *testing.T) {
	request := "LIST osrn:class::* OSP/1.1"

	req, err := osp.Parse(request)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, req.Action, osp.OspActionList)

	assert.Equal(t, req.Osrn, osrn.OSRN{
		Service: "class",
		Type:    "",
		Id:      "*",
	})

	assert.Equal(t, req.Version, "OSP/1.1")
}

func TestParseInvalidAction(t *testing.T) {
	request := "LOL osrn:class::* OSP/1.1"

	_, err := osp.Parse(request)

	if err != nil {
		fmt.Printf("err: %v, target: %v", err.Error(), osp.ErrorBadAction.Error())
		if errors.Is(err, osp.ErrorBadAction) {
			return
		}
		t.Error(err)
	}
}
