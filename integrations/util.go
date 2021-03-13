package integrations

import (
	"fmt"
	"os"
	"time"
)

type Timer struct {
	Name  string
	Start time.Time
}

func NewTimer(name string) *Timer {
	return &Timer{
		Name:  name,
		Start: time.Now(),
	}
}

func (t *Timer) Done() {
	if os.Getenv("HPS1_TIME") != "" {
		_, _ = fmt.Fprintf(os.Stderr, "timer %s: %s\n", t.Name, time.Since(t.Start).String())
	}
}
