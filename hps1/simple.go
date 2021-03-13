package hps1

import (
	"context"
	"fmt"
	"io"
)

type stringComponent string

func (s stringComponent) Render(ctx context.Context, w io.Writer) error {
	_, e := fmt.Fprint(w, s)
	return e
}

func String(s string) Renderer {
	return stringComponent(s)
}

func Empty() Renderer {
	return String("")
}

type stringFuncComponent func(ctx context.Context) (string, error)

func (s stringFuncComponent) Render(ctx context.Context, w io.Writer) error {
	out, e := s(ctx)
	if e != nil {
		return e
	}

	_, e = fmt.Fprint(w, out)
	return e
}

func StringFunc(s func(ctx context.Context) (string, error)) Renderer {
	return stringFuncComponent(s)
}

type Group []Renderer

func (s Group) Render(ctx context.Context, w io.Writer) error {
	for _, comp := range s {
		e := comp.Render(ctx, w)
		if e != nil {
			return e
		}
	}

	return nil
}
