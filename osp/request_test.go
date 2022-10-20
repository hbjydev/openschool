package osp_test

import (
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
	"go.h4n.io/openschool/osp"
	"go.h4n.io/openschool/osrn"
)

func TestOspRequestString(t *testing.T) {
	req := osp.Request{
		Action:  osp.ActionGet,
		Osrn:    osrn.ParseOSRN("osrn:class:classes:*"),
		Version: "OSP/1.1",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: `{"hello":"world"}`,
	}

	request := req.String()
	lines := strings.Split(request, "\r\n")

	assert.Equal(t, len(lines), 4)
	assert.Equal(t, lines[0], "GET osrn:class:classes:* OSP/1.1")
	assert.Equal(t, lines[1], "Content-Type: application/json")
	assert.Equal(t, lines[2], "")
	assert.Equal(t, lines[3], `{"hello":"world"}`)
}
