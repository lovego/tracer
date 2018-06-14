package tracer

import (
	"context"
	"testing"
	"time"
	// "github.com/lovego/xiaomei/utils"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	if got := Context(ctx, nil); got != ctx {
		t.Errorf("unexpected Context(ctx, nil): %v", got)
	}
}

func TestTag(t *testing.T) {
	span := &Span{}
	ctx := Context(context.Background(), span)
	Tag(ctx, "k", "v")
	if len(span.Tags) != 1 || span.Tags["k"] != "v" {
		t.Errorf("unexpected Tags: %v", span.Tags)
	}
}

func TestDebugTag(t *testing.T) {
	span := &Span{debug: true}
	ctx := Context(context.Background(), span)
	DebugTag(ctx, "k", "v")
	if len(span.Tags) != 1 || span.Tags["k"] != "v" {
		t.Errorf("unexpected Tags: %v", span.Tags)
	}
}

func TestGetSpan(t *testing.T) {
	if got := GetSpan(nil); got != nil {
		t.Errorf("unexpected Context(ctx, nil): %v", got)
	}
}

func TestContextDemo(t *testing.T) {
	ctx := context.Background()
	if ctx.Done() != nil {
		t.Error("unexpected non nil Done.")
	}
	ctx, _ = context.WithTimeout(ctx, time.Second)
	if ctx.Done() == nil {
		t.Error("unexpected nil Done.")
	}

	k := struct{}{}
	ctx = context.WithValue(ctx, k, 333)
	if ctx.Done() == nil {
		t.Error("unexpected nil Done.")
	}
	if ctx.Value(k) != 333 {
		t.Error("unexpected value.")
	}
}
