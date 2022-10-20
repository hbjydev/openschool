package osp

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"go.h4n.io/openschool/util"
	"go.uber.org/zap"
)

type Service struct {
	Addr      string
	Name      string
	Resources map[string]Resource

	Tls *tls.Config

	Logger *zap.Logger
}

// Run will start a TCP listener on the network address defined by
// [Service.Addr] and begin a handler loop.
func (s *Service) Run() error {
	var err error
	var lis net.Listener

	// Open a new TCP socket on the [Service.Addr]
	if s.Tls != nil {
		lis, err = tls.Listen("tcp", s.Addr, s.Tls)
		if err != nil {
			return err
		}
	} else {
		lis, err = net.Listen("tcp", s.Addr)
		if err != nil {
			return err
		}
	}

	// Close the listener once the service stops
	defer lis.Close()

	for {
		// Accept the next connection, blocking until one is opened.
		conn, err := lis.Accept()
		if err != nil {
			// If an error is thrown, log it to the console and close the
			// connection.
			s.Logger.Sugar().Error(err)
			os.Exit(1)
		}

		s.Logger.Sugar().Debugw("connection accepted", "remote", conn.RemoteAddr().String())

		// Begin handling the connection on a new thread
		go s.handle(conn)
	}
}

// handle parses and routes an OSP service request, or writes an error back to
// the client if an element of the handler fails.
func (s *Service) handle(conn net.Conn) {
	// Close the connection as soon as the request has been handled (or the
	// method returns).
	defer conn.Close()

	// Read the body of the connection into a string, marking a newline operator
	// as the end of the request body.
	rawRequest, err := util.ReadConnection(conn)
	if err != nil {
		s.Logger.Sugar().Error("failed to read the request", "error", err)
		return
	}

	s.Logger.Sugar().Debugw("read request", "request", rawRequest)

	// Parse the request body as an OSP request.
	req, err := Parse(rawRequest)
	if err != nil {
		resp := Response{
			Status: OspStatusBadRequest,
			Headers: map[string]string{
				"content-type": "text/plain",
			},
			Body: err.Error(),
		}

		_, err := conn.Write(resp.Bytes())
		if err != nil {
			s.Logger.Sugar().Errorf("failed to write response: %v", err.Error())
			return
		}

		return
	}

	logKvs := ConnLogMaps(conn)
	logKvs = append(logKvs, req.LogMaps()...)

	// Ensure the requested service name matches the name of the current
	// service.
	if req.Osrn.Service != s.Name {
		resp := Response{
			Status: OspStatusBadRequest,
			Headers: map[string]string{
				"content-type": "text/plain",
			},
			Body: fmt.Sprintf(`this server does not include service %v`, req.Osrn.Service),
		}

		s.Logger.Sugar().Errorw("invalid service given", logKvs...)

		_, err := conn.Write(resp.Bytes())
		if err != nil {
			s.Logger.Sugar().Errorf("failed to write response: %v", err.Error())
			return
		}

		return
	}

	// Log the request to the console.
	s.Logger.Sugar().Infow("got request", logKvs...)

	res, ok := s.Resources[req.Osrn.Type]
	if !ok {
		resp := Response{
			Status: OspStatusBadRequest,
			Headers: map[string]string{
				"content-type": "text/plain",
			},
			Body: fmt.Sprintf(`this service does not include resource %v`, req.Osrn.Type),
		}

		_, err := conn.Write(resp.Bytes())
		if err != nil {
			s.Logger.Sugar().Errorf("failed to write response: %v", err.Error())
			return
		}

		return
	}

	var resp Response

	switch req.Action {
	case ActionGet:
		resp, err = res.GET(req)
	case ActionList:
		resp, err = res.LIST(req)
	case ActionCreate:
		resp, err = res.CREATE(req)
	case ActionUpdate:
		resp, err = res.UPDATE(req)
	case ActionDelete:
		resp, err = res.DELETE(req)
	}

	if err != nil {
		resp := Response{
			Status: OspStatusServerError,
			Headers: map[string]string{
				"content-type": "text/plain",
			},
			Body: err.Error(),
		}

		_, err := conn.Write(resp.Bytes())
		if err != nil {
			s.Logger.Sugar().Errorf("failed to write response: %v", err.Error())
			return
		}

		return
	}

	_, err = conn.Write(resp.Bytes())
	if err != nil {
		s.Logger.Sugar().Errorf("failed to write response: %v", err.Error())
		return
	}
}

func ConnLogMaps(c net.Conn) []interface{} {
	return []interface{}{
		"conn.remote",
		c.RemoteAddr().String(),
	}
}
