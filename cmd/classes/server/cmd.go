package server

import (
	"github.com/spf13/cobra"
	"go.h4n.io/openschool/class/repos/class"
	"go.h4n.io/openschool/cli"
	"go.h4n.io/openschool/osp"
)

func NewClassesServerCommand() *cobra.Command {
	repo := class.NewInMemoryClassRepository(50)
	classResource := NewClassResource(&repo)

	server := &osp.Service{
		Addr: `0.0.0.0:8001`,
		Name: `classes`,
		Resources: map[string]osp.Resource{
			"class": classResource,
		},
	}

	return cli.CreateCommand(server)
}
