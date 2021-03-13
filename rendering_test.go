package main

import (
	"bytes"
	"context"
	"os"
	"testing"
)

func BenchmarkRendering(t *testing.B) {
	var lines bytes.Buffer
	ctx := context.Background()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < t.N; i++ {
		lines.Reset()

		component, err := ConstructComponents(wd)
		if err != nil {
			t.Fatal(err)
		}

		err = component.Render(ctx, &lines)
		if err != nil {
			t.Fatal(err)
		}
	}
}
