package integrations

import (
	"context"
	"os/exec"
	"strings"
	"unicode"
)

func GitTopLevel(ctx context.Context) (string, error) {
	proc := exec.CommandContext(ctx, "git", "rev-parse", "--show-toplevel")
	proc.Stderr = nil // send to /dev/null
	out, err := proc.Output()
	return strings.TrimRightFunc(string(out), unicode.IsSpace), err
}

func GitSymbolicRef(ctx context.Context) (string, error) {
	proc := exec.CommandContext(ctx, "git", "symbolic-ref", "HEAD")
	proc.Stderr = nil // send to /dev/null
	out, err := proc.Output()

	ref := strings.TrimRightFunc(string(out), unicode.IsSpace)
	ref = strings.TrimPrefix(ref, "refs/heads/")

	return ref, err
}

func GitRevParseShort(ctx context.Context) (string, error) {
	proc := exec.CommandContext(ctx, "git", "rev-parse", "--short", "HEAD")
	proc.Stderr = nil // send to /dev/null
	out, err := proc.Output()
	return strings.TrimRightFunc(string(out), unicode.IsSpace), err
}

func GitStatus(ctx context.Context) (string, error) {
	proc := exec.CommandContext(ctx, "git", "status", "--porcelain", "--ignore-submodules=dirty", "--untracked-files=no")
	proc.Stderr = nil // send to /dev/null
	out, err := proc.Output()
	return strings.TrimRightFunc(string(out), unicode.IsSpace), err
}
