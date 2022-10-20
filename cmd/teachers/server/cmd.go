package server

import (
	"github.com/spf13/cobra"
	"go.h4n.io/openschool/cli"
	"go.h4n.io/openschool/osp"
)

func NewTeachersServerCommand() *cobra.Command {
	server := &osp.Service{
		Addr:      `0.0.0.0:8005`,
		Name:      `teachers`,
		Resources: map[string]osp.Resource{},
	}

	cmd := cli.CreateCommand(server)

	return cmd
}
