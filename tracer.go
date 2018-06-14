package tracer

import (
	"context"
)

var spanKey spanKeyStruct

type Level uint8

type spanKeyStruct struct {
}

func Context(ctx context.Context, s *Span) context.Context {
	if s == nil {
		return ctx
	}
	return context.WithValue(ctx, spanKey, s)
}

func Tag(ctx context.Context, k string, v interface{}) {
	GetSpan(ctx).Tag(k, v)
}

func DebugTag(ctx context.Context, k string, v interface{}) {
	GetSpan(ctx).DebugTag(k, v)
}

func GetSpan(ctx context.Context) *Span {
	if ctx != nil {
		if v := ctx.Value(spanKey); v != nil {
			if s, ok := v.(*Span); ok {
				return s
			}
		}
	}
	return nil
}
