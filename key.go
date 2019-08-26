package tracer

import "context"

var key keyType
var debugKey debugKeyType

type keyType struct{}
type debugKeyType struct{}

func Get(ctx context.Context) *Tracer {
	if ctx != nil {
		if v := ctx.Value(key); v != nil {
			if t, ok := v.(*Tracer); ok {
				return t
			}
		}
	}
	return nil
}

func IsDebug(ctx context.Context) bool {
	if ctx == nil {
		return false
	}
	v := ctx.Value(debugKey)
	b, _ := v.(bool)
	return b
}

func SetDebug(ctx context.Context) context.Context {
	if ctx != nil {
		ctx = context.WithValue(ctx, debugKey, true)
	}
	return ctx
}
