package hps1

type ShellInfo struct {
	ShellName              string
	ShellPromptEscapeStart string
	ShellPromptEscapeEnd   string
}

var ShellBash = &ShellInfo{ShellName: "bash", ShellPromptEscapeStart: "\\[", ShellPromptEscapeEnd: "\\]"}

var ShellZsh = &ShellInfo{ShellName: "zsh", ShellPromptEscapeStart: "%{", ShellPromptEscapeEnd: "%}"}

var ShellUnknown = &ShellInfo{}
