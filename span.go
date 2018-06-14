package tracer

import (
	"context"
	"time"
)

// Span for tracing
type Span struct {
	Name     string                 `json:"name,omitempty"`
	At       time.Time              `json:"at"`
	Duration float64                `json:"duration"` // milliseconds
	Children []*Span                `json:"children,omitempty"`
	Tags     map[string]interface{} `json:"tags,omitempty"`
	debug    bool
}

// StartSpan start a new span on the given context
func StartSpan(ctx context.Context, name string) *Span {
	if ctx != nil {
		if value := ctx.Value(spanKey); value != nil {
			if parent, ok := value.(*Span); ok && parent != nil {
				s := &Span{Name: name, At: time.Now()}
				parent.Children = append(parent.Children, s)
				return s
			}
		}
	}
	return nil
}

// Finish a span
func (s *Span) Finish() {
	if s != nil {
		s.Duration = float64(time.Since(s.At)) / float64(time.Millisecond)
	}
}

// Tag add a tag to a span
func (s *Span) Tag(k string, v interface{}) *Span {
	if s == nil {
		return s
	}
	if s.Tags == nil {
		s.Tags = make(map[string]interface{})
	}
	s.Tags[k] = v
	return s
}

// Debug add a tag to a span
func (s *Span) Debug() bool {
	if s != nil {
		return s.debug
	}
	return false
}

// SetDebug set if debugging tag should be traced.
func (s *Span) SetDebug(b bool) *Span {
	if s != nil {
		s.debug = b
	}
	return s
}

// DebugTag add a debug tag to a span
func (s *Span) DebugTag(k string, v interface{}) *Span {
	if s == nil || !s.debug {
		return s
	}
	if s.Tags == nil {
		s.Tags = make(map[string]interface{})
	}
	s.Tags[k] = v
	return s
}
