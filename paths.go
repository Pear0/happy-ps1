package hps1

import (
	"path/filepath"
	"strings"
)

func PrettifyPath(absPath string, minComponents int, maxLength int) string {
	components := strings.Split(absPath, string(filepath.Separator))

	pathLength := len(absPath)
	fullIndex := 0

	// as much as possible, contract parent path lengths, to try to go below max length
	for fullIndex < len(components)-minComponents && pathLength > maxLength {
		if len(components[fullIndex]) > 0 {
			pathLength -= len(components[fullIndex]) - 1
		}
		fullIndex++
	}

	for i := 0; i < fullIndex; i++ {
		if len(components[i]) > 1 {
			components[i] = components[i][:1]
		}
	}

	combined := filepath.Join(components...)
	if len(components) > 0 && components[0] == "" {
		combined = string(filepath.Separator) + combined
	}

	return combined
}
