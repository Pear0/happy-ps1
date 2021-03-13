package hps1

import (
	"context"
	"fmt"
	"github.com/Pear0/happy-ps1/integrations"
	"os"
	"strings"
)

func User() Renderer {
	return String(os.Getenv("USER"))
}

func Hostname() Renderer {
	return StringFunc(func(ctx context.Context) (string, error) {
		defer integrations.NewTimer("hostname -s").Done()

		out, err := os.Hostname()
		if err != nil {
			return "", err
		}

		domains := strings.Split(out, ".")

		return domains[0], err
	})
}

func GitPromptInfo() Renderer {
	return StringFunc(func(ctx context.Context) (string, error) {

		// no git here, bail early
		root, _ := integrations.GitIsInWorktree(ctx)
		if root == "" {
			return "", nil
		}

		branchChan := make(chan string)
		refChan := make(chan string)
		stateChan := make(chan string)

		go func() {
			branch, _ := integrations.GitSymbolicRef(ctx)
			branchChan <- branch
		}()

		go func() {
			ref, _ := integrations.GitRevParseShort(ctx)
			refChan <- ref
		}()

		go func() {
			var state string
			if out, err := integrations.GitStatus(ctx); err == nil {
				if len(strings.TrimSpace(out)) > 0 {
					state = "*"
				}
			} else {
				state = "?"
			}

			stateChan <- state
		}()

		branch := <-branchChan
		ref := <-refChan

		// cant find the HEAD info
		if branch == "" && ref == "" {
			return "", nil
		}

		if branch == "" {
			branch = ref
		}

		state := <-stateChan

		return fmt.Sprintf("‹%s%s›", branch, state), nil
	})
}
