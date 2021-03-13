package integrations

import (
	"context"
	"os"
	"os/exec"
)

func Exec(ctx context.Context, args ...string) *exec.Cmd {
	proc := exec.CommandContext(ctx, args[0], args[1:]...)
	proc.Stderr = os.Stderr
	return proc
}
