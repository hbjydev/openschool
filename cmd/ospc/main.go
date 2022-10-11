package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"go.h4n.io/openschool/osp"
	"go.h4n.io/openschool/osrn"
	"go.h4n.io/openschool/util"
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

	req := osp.OspRequest{
		Action:  action,
		Osrn:    osrn.ParseOSRN(strings.ToLower(os.Args[3])),
		Version: "OSP/1.1",
	}

	conn.Write(req.Bytes())

	lines, err := util.ReadConnection(conn)
	if err != nil {
		fmt.Printf("error reading response: %v", err.Error())
		return
	}

	fmt.Println(lines)
}
