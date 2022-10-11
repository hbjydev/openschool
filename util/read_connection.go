package util

import (
	"bufio"
	"fmt"
	"net"
)

func ReadConnection(conn net.Conn) (string, error) {
	data := bufio.NewScanner(conn)
	lines := ``

	for data.Scan() {
		line := data.Text()
		lines = fmt.Sprintf("%v\n", line)
	}

	if data.Err() != nil {
		return "", data.Err()
	}

	return lines, nil
}
