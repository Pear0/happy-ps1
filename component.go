package hps1

import (
	"bytes"
	"context"
	"io"
)

type Renderer interface {
	Render(ctx context.Context, w io.Writer) error
}

type Composer interface {
	Compose(r Renderer) Renderer
}

func RenderToString(ctx context.Context, r Renderer) (string, error) {
	var buf bytes.Buffer

	err := r.Render(ctx, &buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
