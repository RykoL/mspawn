package internal

import (
	"context"
	"os"
	"os/exec"
)

func SpawnProcess(ctx context.Context, cmd string, args []string) error {

	process := exec.CommandContext(ctx, cmd, args...)
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	if err := process.Run(); err != nil {
		return err
	}
	return nil
}
