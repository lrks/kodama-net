package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	commit  = ""
	buildAt = ""
)

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "commit:  %s\nbuildAt: %s\n", commit, buildAt)
			return fmt.Errorf("failed to print version information: %w", err)
		},
	}

	return cmd
}
