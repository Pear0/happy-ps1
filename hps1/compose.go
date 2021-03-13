package hps1

import (
	"context"
	"github.com/fatih/color"
	"io"
)

type appliedColor struct {
	color *color.Color
	r     Renderer
}

func (s appliedColor) Render(ctx context.Context, w io.Writer) error {

	out, err := RenderToString(ctx, s.r)
	if err != nil {
		return err
	}

	if len(out) == 0 {
		return nil
	}

	_, e := s.color.Fprint(w, out)
	return e
}

type colorComponent color.Color

func (s *colorComponent) Compose(r Renderer) Renderer {
	return appliedColor{color: (*color.Color)(s), r: r}
}

func Color(c *color.Color) Composer {
	return (*colorComponent)(c)
}
