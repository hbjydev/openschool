package osp

import (
	"fmt"
	"strings"
)

type OspStatus struct {
	StatusCode int
	Reason     string
}

func (s *OspStatus) String() string {
	return fmt.Sprintf("%v %v", s.StatusCode, s.Reason)
}

var (
	OspStatusSuccess = OspStatus{
		StatusCode: 400,
		Reason:     "Bad Request",
	}
	OspStatusBadRequest = OspStatus{
		StatusCode: 400,
		Reason:     "Bad Request",
	}
)

type Response struct {
	Version string
	Status  OspStatus
	Headers map[string]string
	Body    string
}

func (r *Response) String() string {
	if r.Version == "" {
		r.Version = "OSP/1.1"
	}

	str := fmt.Sprintf("%v %v", r.Version, r.Status.String())

	if len(r.Headers) > 0 {
		var headersList []string
		for k, v := range r.Headers {
			line := fmt.Sprintf("%v: %v", k, v)
			headersList = append(headersList, line)
		}
		str = fmt.Sprintf("%v\n%v", str, strings.Join(headersList, "\n"))
	}

	if r.Body != "" {
		str = fmt.Sprintf("%v\n%v", str, r.Body)
	}

	return str
}

func (r *Response) Bytes() []byte {
	resp := r.String()
	return []byte(resp)
}