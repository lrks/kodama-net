package main

import (
	"github.com/lrks/kodama-net/internal/discovery"
	"github.com/lrks/kodama-net/internal/hello"
	"github.com/spf13/cobra"
)

type container struct {
	helloSvc     hello.Service
	discoverySvc discovery.Service
}

func (c *container) HelloService() hello.Service {
	return c.helloSvc
}

func (c *container) DiscoveryService() discovery.Service {
	return c.discoverySvc
}

func newRootCmd() *cobra.Command {
	container := &container{
		helloSvc:     hello.NewService(),
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
		newHelloCmd(container),
		newDiscoveryCmd(container),
	)

	return cmd
}
