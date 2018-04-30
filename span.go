package tracer

import (
	"context"
	"time"
)

type spanKeyStruct struct {
}

var spanKey spanKeyStruct

func SpanKey() spanKeyStruct {
	return spanKey
}

type Span struct {
	Name     string
	At       time.Time
	Duration float64 // milliseconds
	Children []*Span
}

func (s *Span) Finish() {
	s.Duration = float64(time.Since(s.At)) / float64(time.Millisecond)
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
