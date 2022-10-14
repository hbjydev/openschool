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

	logger *zap.Logger
}

// Run will start a TCP listener on the network address defined by
// [Service.Addr] and begin a handler loop.
func (s *Service) Run() error {
	// Create a new Zap logger (which will log to JSON)
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	s.logger = logger

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
			logger.Sugar().Error(err)
			os.Exit(1)
		}

		s.logger.Sugar().Debugw("connection accepted", "remote", conn.RemoteAddr().String())

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
		s.logger.Sugar().Error("failed to read the request", "error", err)
		return
	}

	s.logger.Sugar().Debugw("read request", "request", rawRequest)

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
		conn.Write(resp.Bytes())
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

		s.logger.Sugar().Errorw("invalid service given", logKvs...)

		conn.Write(resp.Bytes())
		return
	}

	// Log the request to the console.
	s.logger.Sugar().Infow("got request", logKvs...)

	res, ok := s.Resources[req.Osrn.Type]
	if !ok {
		resp := Response{
			Status: OspStatusBadRequest,
			Headers: map[string]string{
				"content-type": "text/plain",
			},
			Body: fmt.Sprintf(`this service does not include resource %v`, req.Osrn.Type),
		}

		conn.Write(resp.Bytes())
		return
	}

	var resp Response

	switch req.Action {
	case OspActionGet:
		resp, err = res.GET(req)
	case OspActionList:
		resp, err = res.LIST(req)
	case OspActionCreate:
		resp, err = res.CREATE(req)
	case OspActionUpdate:
		resp, err = res.UPDATE(req)
	case OspActionDelete:
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

		conn.Write(resp.Bytes())
		return
	}

	conn.Write(resp.Bytes())
}

func ConnLogMaps(c net.Conn) []interface{} {
	return []interface{}{
		"conn.remote",
		c.RemoteAddr().String(),
	}
}
