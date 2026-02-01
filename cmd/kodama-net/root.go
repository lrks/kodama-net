package main

import (
	"github.com/lrks/kodama-net/internal/discovery"
	"github.com/spf13/cobra"
)

type container struct {
	discoverySvc discovery.Service
}

func (c *container) DiscoveryService() discovery.Service {
	return c.discoverySvc
}

func newRootCmd() *cobra.Command {
	container := &container{
		discoverySvc: discovery.NewService(),
	}

	cmd := &cobra.Command{
		Use:   "kodama-net",
		Short: "kodama-net: ECHONET Lite explorer",
		RunE: func(cmd *cobra.Command, args []string) error {
			// ルート直下では特に処理を行わず、サブコマンドに委譲する
			return nil
		},
	}

	// サブコマンド
	cmd.AddCommand(
		newVersionCmd(),
		newDiscoveryCmd(container),
	)

	return cmd
}
