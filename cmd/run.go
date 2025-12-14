package cmd

import (
	"context"
	"os"
	"os/signal"
	"strings"

	"github.com/rykol/mspawn/internal"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var commands []string

func splitCmd(cmd string) (string, []string) {
	splittedCmd := strings.Fields(cmd)

	if len(splittedCmd) == 1 {
		return cmd, []string{}
	}

	program := splittedCmd[0]
	args := splittedCmd[1:]
	return program, args
}

var runCmd = &cobra.Command{
	Use:   "run --cmd cmd1",
	Short: "Run multiple processes in paralell",
	RunE: func(cmd *cobra.Command, args []string) error {

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
		defer stop()

		errorGroup, groupCtx := errgroup.WithContext(ctx)

		for _, command := range commands {
			cmd, args := splitCmd(command)
			errorGroup.Go(func() error {
				return internal.SpawnProcess(groupCtx, cmd, args)
			})

		}

		if err := errorGroup.Wait(); err != nil {
			if ctx.Err() != nil {
				return nil
			}
			return err
		}

		return nil
	},
}
