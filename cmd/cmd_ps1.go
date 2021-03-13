package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Pear0/hps1"
	"github.com/Pear0/hps1/integrations"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"math"
	"os"
)

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
	fmt.Fprintf(os.Stderr, "happy-ps1 error: %s\n", err.Error())
	fmt.Printf(" $")
	os.Exit(1)
}

func warnPs1Error(err error) {
	fmt.Fprintf(os.Stderr, "happy-ps1 error: %s\n", err.Error())
}

func handlePs1Command(cmd *cobra.Command, args []string) {
	var lines bytes.Buffer

	ctx := context.Background()

	component, err := constructComponents()
	if err != nil {
		fatalPs1Error(err)
	}

	err = component.Render(ctx, &lines)
	if err != nil {
		fatalPs1Error(err)
	}

	_, _ = io.Copy(os.Stdout, &lines)
}

func constructComponents() (hps1.Renderer, error) {
	wd, err := os.Getwd()
	if err != nil {
		// this is a critical error
		return nil, err
	}

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
		hps1.Color(color.New(color.FgRed)).Compose(hps1.Group{
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
		hps1.String("$ "),
	}, nil
}
