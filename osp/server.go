package osp

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"go.uber.org/zap"
)

type Service struct {
	Addr string
	Name string

	logger *zap.Logger
}

// Run will start a TCP listener on the network address defined by
// [Service.Addr] and begin a handler loop.
func (s *Service) Run() error {
	// Create a new Zap logger (which will log to JSON)
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	s.logger = logger

	// Open a new TCP socket on the [Service.Addr]
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
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
	rawRequest, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		s.logger.Sugar().Error("failed to read the request", "error", err)
		return
	}

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
}

func ConnLogMaps(c net.Conn) []interface{} {
	return []interface{}{
		"conn.remote",
		c.RemoteAddr().String(),
	}
}
