package tracer

import (
	"context"
	"time"
)

type spanKeyStruct struct {
}

var spanKey spanKeyStruct

type Span struct {
	Name     string                 `json:"name,omitempty"`
	At       time.Time              `json:"at"`
	Duration float64                `json:"duration"` // milliseconds
	Children []*Span                `json:"children"`
	Tags     map[string]interface{} `json:"tags"`
}

func StartSpan(ctx context.Context, name string) *Span {
	s := &Span{Name: name, At: time.Now()}
	if ctx != nil {
		if value := ctx.Value(spanKey); value != nil {
			if parent, ok := value.(*Span); ok && parent != nil {
				parent.Children = append(parent.Children, s)
			}
		}
	}
	return s
}

func (s *Span) Finish() {
	s.Duration = float64(time.Since(s.At)) / float64(time.Millisecond)
}

func (s *Span) Tag(k string, v interface{}) {
	if s == nil {
		return
	}
	if s.Tags == nil {
		s.Tags = make(map[string]interface{})
	}
	s.Tags[k] = v
}

func Context(ctx context.Context, s *Span) context.Context {
	return context.WithValue(ctx, spanKey, s)
}
func Tag(ctx context.Context, k string, v interface{}) {
	GetSpan(ctx).Tag(k, v)
}

func GetSpan(ctx context.Context) *Span {
	if v := ctx.Value(spanKey); v != nil {
		if s, ok := v.(*Span); ok {
			return s
		}
	}
	return nil
}
