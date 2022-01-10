package hps1

import "context"

type contextKey int

var keyShellId contextKey = 1

func WrapContextShellInfo(ctx context.Context, info *ShellInfo) context.Context {
	return context.WithValue(ctx, keyShellId, info)
}

func GetShellInfo(ctx context.Context) *ShellInfo {
	v := ctx.Value(keyShellId)
	if v != nil {
		return v.(*ShellInfo)
	} else {
		return nil
	}
}
