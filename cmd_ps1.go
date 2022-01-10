package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Pear0/happy-ps1/hps1"
	"github.com/Pear0/happy-ps1/integrations"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"math"
	"os"
	"path"
	"time"
)

var programStartTime = time.Now()

var ps1LastExitCode int

var ps1Cmd = &cobra.Command{
	Use:   "ps1",
	Short: "generate ps1 line",
	Run:   handlePs1Command,
}

func init() {
	rootCmd.AddCommand(ps1Cmd)

	ps1Cmd.Flags().IntVar(&ps1LastExitCode, "last-exit", math.MaxInt32, "exit code of the prior command")
}

func GetLastExitCode() (int, bool) {
	if ps1LastExitCode == math.MaxInt32 {
		return -1, false
	}
	return ps1LastExitCode, true
}

func fatalPs1Error(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "happy-ps1 error: %s\n", err.Error())
	fmt.Printf(" $")
	os.Exit(1)
}

func warnPs1Error(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "happy-ps1 error: %s\n", err.Error())
}

func createShellInfo() *hps1.ShellInfo {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return hps1.ShellUnknown
	}

	shellName := path.Base(shell)

	if shellName == "zsh" {
		return hps1.ShellZsh
	} else if shellName == "bash" {
		return hps1.ShellBash
	} else {
		warnPs1Error(errors.Errorf("unknown $SHELL %s -> %s\n", shell, shellName))
		return hps1.ShellUnknown
	}
}

func handlePs1Command(cmd *cobra.Command, args []string) {
	var lines bytes.Buffer

	ctx := context.Background()
	ctx = hps1.WrapContextShellInfo(ctx, createShellInfo())

	wd, err := os.Getwd()
	if err != nil {
		fatalPs1Error(err)
	}

	component, err := ConstructComponents(wd)
	if err != nil {
		fatalPs1Error(err)
	}

	err = component.Render(ctx, &lines)
	if err != nil {
		fatalPs1Error(err)
	}

	_, _ = io.Copy(os.Stdout, &lines)
}

func ConstructComponents(wd string) (hps1.Renderer, error) {
	mnt, err := integrations.GetMountForPath(wd)
	if err != nil {
		warnPs1Error(err)
	}

	return hps1.Group{
		hps1.Color(color.New(color.FgRed)).Compose(hps1.Group{
			func() hps1.Renderer {
				if code, ok := GetLastExitCode(); ok && code != 0 {
					return hps1.Group{
						hps1.String(fmt.Sprintf("[%d] ", code)),
					}
				} else {
					return hps1.Empty()
				}
			}(),
		}),
		hps1.Color(color.New(color.FgGreen)).Compose(hps1.Group{
			hps1.User(),
			hps1.String("@"),
			hps1.Hostname(),
			hps1.String(":"),
		}),
		hps1.Color(color.New(color.FgHiBlue, color.Bold)).Compose(hps1.Group{
			hps1.String(hps1.PrettifyPath(wd, 2, 25)),
		}),
		hps1.Color(color.New(color.FgHiRed, color.Bold)).Compose(hps1.Group{
			func() hps1.Renderer {
				if !mnt.IsRemote {
					return hps1.Group{
						hps1.String(" "),
						hps1.GitPromptInfo(),
					}
				} else {
					return hps1.Empty()
				}
			}(),
		}),
		hps1.StringFunc(func(ctx context.Context) (string, error) {
			if os.Getenv("HPS1_TIME") != "" {
				return fmt.Sprintf("(%s)", time.Since(programStartTime).String()), nil
			} else {
				return "", nil
			}
		}),
		hps1.String("$ "),
	}, nil
}
