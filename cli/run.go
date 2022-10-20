package cli

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func Run(service string, cmd *cobra.Command) int {
	cmd.Short = service

	if err := run(cmd); err != nil {
		return 1
	}

	return 0
}

func run(cmd *cobra.Command) (err error) {
	rand.Seed(time.Now().UnixNano())

	if !cmd.SilenceUsage {
		cmd.SilenceUsage = true
		cmd.SetFlagErrorFunc(func(c *cobra.Command, err error) error {
			c.SilenceUsage = false
			return err
		})
	}

	cmd.SilenceErrors = true

	switch {
	case cmd.PersistentPreRun != nil:
		pre := cmd.PersistentPreRun
		cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
			pre(cmd, args)
		}

	case cmd.PersistentPreRunE != nil:
		pre := cmd.PersistentPreRunE
		cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			return pre(cmd, args)
		}

	default:
		cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {}
	}

	err = cmd.Execute()
	return
}
