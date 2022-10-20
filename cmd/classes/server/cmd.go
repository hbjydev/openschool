package server

import (
	"crypto/tls"
	"errors"
	"os"

	"github.com/spf13/cobra"
	"go.h4n.io/openschool/class/repos/class"
	"go.h4n.io/openschool/osp"
	"go.uber.org/zap"
)

func NewClassesServerCommand() *cobra.Command {
	cmd := cobra.Command{
		Short: `classes`,
		Use:   `classes [-L addr] [--tls.enable] [--tls.chain path/to/chain] [--tls.key path/to/key]`,

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
			repo := class.NewInMemoryClassRepository(50)
			classResource := NewClassResource(&repo)

			server := &osp.Service{
				Addr: `0.0.0.0:8001`,
				Name: `classes`,
				Resources: map[string]osp.Resource{
					"class": classResource,
				},
				Logger: logger,
			}

			if tlsEnabled {
				server.Logger.Sugar().Info("server configured to use TLS, configuring...")
				cert, err := tls.LoadX509KeyPair("pki/classes/fullchain.pem", "pki/classes/cert.key")
				if err != nil {
					panic(err)
				}
				cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
				server.Tls = cfg
			}

			server.Logger.Sugar().Info("starting server", "addr", server.Addr)
			return server.Run()
		},
	}

	cmd.PersistentFlags().StringP("listen", "L", "0.0.0.0:8001", "the address to listen for incoming connections on")
	cmd.PersistentFlags().Bool("tls.enable", false, "whether or not to enable TLS")
	cmd.PersistentFlags().String("tls.chain", "/etc/openschool/classes/tls.crt", "the TLS full-chain PEM to configure CAs and certificate data")
	cmd.PersistentFlags().String("tls.key", "/etc/openschool/classes/tls.key", "the TLS private key to authenticate the server with")

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
