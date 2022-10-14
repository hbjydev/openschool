package util

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func ReadConnection(conn net.Conn) (string, error) {
	data := bufio.NewScanner(conn)
	lines := ``

	data.Split(bufio.ScanLines)

	for data.Scan() {
		if len(data.Text()) == 0 {
			break
		}

		log.Println(data.Text())
		lines = fmt.Sprintf("%v", data.Text())
	}

	if data.Err() != nil {
		return "", data.Err()
	}

	return lines, nil
}
