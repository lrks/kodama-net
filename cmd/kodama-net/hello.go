package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func newHelloCmd(container *container) *cobra.Command {
	var name string

	svc := container.HelloService()

	cmd := &cobra.Command{
		Use:   "hello",
		Short: "Output a greeting message",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			msg, err := svc.Greet(ctx, name)
			if err != nil {
				return fmt.Errorf("failed to greet: %w", err)
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), msg)
			if err != nil {
				return fmt.Errorf("failed to write output: %w", err)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "world", "Target name for greeting (default: world)")

	return cmd
}
