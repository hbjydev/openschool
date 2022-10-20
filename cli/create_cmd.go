package cli

import (
	"crypto/tls"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.h4n.io/openschool/osp"
	"go.uber.org/zap"
)

func CreateCommand(server *osp.Service) *cobra.Command {
	use := fmt.Sprintf("%v [-L addr] [--tls.enable] [--tls.chain path/to/chain] [--tls.key path/to/key]", server.Name)

	cmd := cobra.Command{
		Short: server.Name,
		Use:   use,

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			tlsEnabled, err := cmd.PersistentFlags().GetBool("tls.enable")
			if err != nil {
				return err
			}

			if tlsEnabled {
				validateTlsFiles(cmd)
			}

			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			logger, err := zap.NewDevelopment()
			if err != nil {
				return err
			}

			tlsEnabled, err := cmd.PersistentFlags().GetBool("tls.enable")
			if err != nil {
				return err
			}

			addr, err := cmd.PersistentFlags().GetString("listen")
			if err != nil {
				return err
			}

			server.Addr = addr
			server.Logger = logger

			if tlsEnabled {
				chain, err := cmd.PersistentFlags().GetString("tls.chain")
				if err != nil {
					return err
				}

				key, err := cmd.PersistentFlags().GetString("tls.key")
				if err != nil {
					return err
				}

				server.Logger.Sugar().Infow("server configured to use TLS, configuring...")
				cert, err := tls.LoadX509KeyPair(chain, key)
				if err != nil {
					panic(err)
				}
				cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
				server.Tls = cfg
			}

			server.Logger.Sugar().Infow("starting server", "addr", server.Addr, "name", server.Name)
			return server.Run()
		},
	}

	cmd.PersistentFlags().StringP("listen", "L", server.Addr, "the address to listen for incoming connections on")
	cmd.PersistentFlags().Bool("tls.enable", false, "whether or not to enable TLS")
	cmd.PersistentFlags().String("tls.chain", fmt.Sprintf("/etc/openschool/%v/tls.crt", server.Name), "the TLS full-chain PEM to configure CAs and certificate data")
	cmd.PersistentFlags().String("tls.key", fmt.Sprintf("/etc/openschool/%v/tls.key", server.Name), "the TLS private key to authenticate the server with")

	return &cmd
}

func validateTlsFiles(cmd *cobra.Command) error {
	chain, err := cmd.PersistentFlags().GetString("tls.chain")
	if err != nil {
		return err
	}

	key, err := cmd.PersistentFlags().GetString("tls.key")
	if err != nil {
		return err
	}

	chainInfo, err := os.Stat(chain)
	if err != nil {
		return err
	}

	if chainInfo.IsDir() {
		return errors.New("provided tls chain is a directory, not a file")
	}

	keyInfo, err := os.Stat(key)
	if err != nil {
		return err
	}

	if keyInfo.IsDir() {
		return errors.New("provided tls key is a directory, not a file")
	}

	return nil
}
