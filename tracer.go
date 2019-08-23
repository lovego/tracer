package tracer

import (
	"context"
	"fmt"
	"time"
)

type tracer struct {
	Name     string                 `json:"name,omitempty"`
	At       time.Time              `json:"at"`
	Duration float64                `json:"duration"` // milliseconds
	Children []*tracer              `json:"children,omitempty"`
	Tags     map[string]interface{} `json:"tags,omitempty"`
	Logs     []string               `json:"logs,omitempty"`
}

// Start start a new tracer on the given context
func Start(ctx context.Context, name string) context.Context {
	tracer := tracer{Name: name, At: time.Now()}

	if parent := get(ctx); parent != nil {
		parent.Children = append(parent.Children, tracer)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, key, tracer)
}

// Finish finish the tracer on the given context
func Finish(ctx context.Context) {
	if tracer := get(ctx); tracer != nil {
		tracer.Duration = float64(time.Since(tracer.At)) / float64(time.Millisecond)
	}
}

// Tag add a tag to a tracer context
func Tag(ctx context.Context, k string, v interface{}) {
	if tracer := get(ctx); tracer != nil {
		if tracer.Tags == nil {
			tracer.Tags = make(map[string]interface{})
		}
		tracer.Tags[k] = v
	}
}

// DebugTag add a tag to a tracer context if debug is enabled
func DebugTag(ctx context.Context, k string, v interface{}) {
	if ctx != nil && ctx.Value("debug") != nil {
		Tag(ctx, k, v)
	}
}

// Log add a log to a tracer context using Sprint
func Log(ctx context.Context, args ...interface{}) {
	if tracer := get(ctx); tracer != nil {
		s.Logs = append(s.Logs, fmt.Sprint(args...))
	}
}

// Logf add a log to a tracer context using Sprintf
func Logf(ctx context.Context, format string, args ...interface{}) {
	if tracer := get(ctx); tracer != nil {
		s.Logs = append(s.Logs, fmt.Sprintf(format, args...))
	}
}

// DebugLog add a log to a tracer context if debug is enabled
func DebugLog(ctx context.Context, args ...interface{}) {
	GetSpan(ctx).DebugLog(args...)
}

// DebugLogf add a log to a tracer context if debug is enabled
func DebugLogf(ctx context.Context, format string, args ...interface{}) {
	GetSpan(ctx).DebugLogf(format, args...)
}

var key keyType

type keyType struct{}

func get(ctx context.Context) *tracer {
	if ctx != nil {
		if v := ctx.Value(key); v != nil {
			if t, ok := v.(*tracer); ok {
				return t
			}
		}
	}
	return nil
}
