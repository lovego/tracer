package tracer

import (
	"context"
)

var spanKey spanKeyStruct

type spanKeyStruct struct {
}

// Context return a tracer context with the span
func Context(ctx context.Context, s *Span) context.Context {
	if s == nil {
		return ctx
	}
	return context.WithValue(ctx, spanKey, s)
}

// StartContext start a new context on the given context
func StartContext(ctx context.Context, name string) context.Context {
	return context.WithValue(context.Background(), spanKey, StartSpan(ctx, name))
}

// Tag add a tag to tracer context
func Tag(ctx context.Context, k string, v interface{}) {
	GetSpan(ctx).Tag(k, v)
}

// DebugTag add a debug tag to tracer context
func DebugTag(ctx context.Context, k string, v interface{}) {
	GetSpan(ctx).DebugTag(k, v)
}

// GetSpan get span from a tracer context
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
