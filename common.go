package hps1

import (
	"context"
	"fmt"
	"github.com/Pear0/hps1/integrations"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

func User() Renderer {
	return String(os.Getenv("USER"))
}

func Hostname() Renderer {
	return StringFunc(func(ctx context.Context) (string, error) {
		proc := exec.CommandContext(ctx, "hostname", "-s")
		proc.Stderr = os.Stderr
		out, err := proc.Output()
		return strings.TrimRightFunc(string(out), unicode.IsSpace), err
	})
}

func GitPromptInfo() Renderer {
	return StringFunc(func(ctx context.Context) (string, error) {

		branch, _ := integrations.GitSymbolicRef(ctx)
		ref, _ := integrations.GitRevParseShort(ctx)

		// no git here
		if branch == "" && ref == "" {
			return "", nil
		}

		if branch == "" {
			branch = ref
		}

		var state string

		if out, err := integrations.GitStatus(ctx); err == nil {
			if len(strings.TrimSpace(out)) > 0 {
				state = "*"
			}
		} else {
			state = "?"
		}

		return fmt.Sprintf("<%s%s>", branch, state), nil
	})
}
