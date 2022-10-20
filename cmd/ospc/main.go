package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"go.h4n.io/openschool/osp"
	"go.h4n.io/openschool/osrn"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("no endpoint, action, or osrn defined")
		return
	}

	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	actionRaw := strings.ToUpper(os.Args[2])
	action, ok := osp.GetAction(actionRaw)
	if !ok {
		fmt.Printf("invalid osp action: %v", actionRaw)
		return
	}

	req := osp.Request{
		Action:  action,
		Osrn:    osrn.ParseOSRN(strings.ToLower(os.Args[3])),
		Version: "OSP/1.1",
	}

	fmt.Printf("%v\n", req.String())
	reqData := []byte(fmt.Sprintf("%v\r\n\r\n", req.String()))

	_, err = conn.Write(reqData)
	if err != nil {
		log.Fatalf("failed to write request: %v", err.Error())
		return
	}

	result, err := io.ReadAll(conn)
	if err != nil {
		fmt.Printf("error reading response: %v", err.Error())
	}
	fmt.Println(string(result))
}
