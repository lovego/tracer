package tracer

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Tracer struct {
	sync.Mutex

	Name     string                 `json:"name,omitempty"`
	At       time.Time              `json:"at"`
	Duration float64                `json:"duration"` // milliseconds
	Children []*Tracer              `json:"children,omitempty"`
	Tags     map[string]interface{} `json:"tags,omitempty"`
	Logs     []string               `json:"logs,omitempty"`
}

// Start start a new tracer on the given context
func Start(ctx context.Context, name string) context.Context {
	if ctx == nil {
		return nil
	}
	tracer := &Tracer{Name: name, At: time.Now()}

	if parent := Get(ctx); parent != nil {
		parent.Children = append(parent.Children, tracer)
	}
	return context.WithValue(ctx, key, tracer)
}

// Start start a child tracer on the given context
func StartChild(ctx context.Context, name string) context.Context {
	parent := Get(ctx)
	if parent == nil {
		return ctx
	}
	tracer := &Tracer{Name: name, At: time.Now()}
	parent.Children = append(parent.Children, tracer)
	return context.WithValue(ctx, key, tracer)
}

// Finish finish the tracer on the given context
func Finish(ctx context.Context) {
	if tracer := Get(ctx); tracer != nil {
		tracer.Duration = float64(time.Since(tracer.At)) / float64(time.Millisecond)
	}
}

// Tag add a tag to a tracer context
func Tag(ctx context.Context, k string, v interface{}) {
	if tracer := Get(ctx); tracer != nil {
		if tracer.Tags == nil {
			tracer.Tags = make(map[string]interface{})
		}
		tracer.Lock()
		tracer.Tags[k] = v
		tracer.Unlock()
	}
}

// TagString a tag Marshal to string and add to a tracer context
func TagString(ctx context.Context, k string, v interface{}) {
	if tracer := Get(ctx); tracer != nil {
		if tracer.Tags == nil {
			tracer.Tags = make(map[string]interface{})
		}
		b, _ := json.Marshal(v)
		tracer.Lock()
		tracer.Tags[k] = string(b)
		tracer.Unlock()
	}
}

// DebugTag add a tag to a tracer context if debug is enabled
func DebugTag(ctx context.Context, k string, v interface{}) {
	if IsDebug(ctx) {
		Tag(ctx, k, v)
	}
}

// Log add a log to a tracer context using Sprint
func Log(ctx context.Context, args ...interface{}) {
	if tracer := Get(ctx); tracer != nil {
		tracer.Logs = append(tracer.Logs, fmt.Sprint(args...))
	}
}

// Logf add a log to a tracer context using Sprintf
func Logf(ctx context.Context, format string, args ...interface{}) {
	if tracer := Get(ctx); tracer != nil {
		tracer.Logs = append(tracer.Logs, fmt.Sprintf(format, args...))
	}
}

// DebugLog add a log to a tracer context if debug is enabled
func DebugLog(ctx context.Context, args ...interface{}) {
	if IsDebug(ctx) {
		Log(ctx, args...)
	}
}

// DebugLogf add a log to a tracer context if debug is enabled
func DebugLogf(ctx context.Context, format string, args ...interface{}) {
	if IsDebug(ctx) {
		Logf(ctx, format, args...)
	}
}
