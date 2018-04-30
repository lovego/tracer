package tracer

import (
	"context"
	"fmt"
	"testing"
	// "github.com/lovego/xiaomei/utils"
)

func TestStartSpan(t *testing.T) {
	var span *Span
	func() {
		ctx := context.Background()
		span = StartSpan(ctx, "root")
		defer span.Finish()
		ctx = context.WithValue(ctx, SpanKey(), span)

		testRunSpan(ctx, "task0")
		testRunSpan(ctx, "task1", "t1.0", "t1.1", "t1.2")
		testRunSpan(ctx, "task2", "t2.0", "t2.1", "t2.2")
	}()
	// utils.PrintJson(span)
	if span.Name != "root" {
		t.Errorf("unexpected name: %v", span.Name)
	}
	if span.Duration <= 0 {
		t.Errorf("unexpected duration: %v", span.Duration)
	}
	if len(span.Children) != 3 {
		t.Errorf("unexpected Children length: %d", len(span.Children))
	}
	for i, children := range []int{0, 3, 3} {
		if span.Children[i].Name != fmt.Sprintf("task%d", i) {
			t.Errorf("unexpected task%d name: %v", i, span.Children[i].Name)
		}
		if span.Children[i].Duration <= 0 {
			t.Errorf("unexpected task%d duration: %v", i, span.Children[i].Duration)
		}
		if len(span.Children[i].Children) != children {
			t.Errorf("unexpected task%d Children length: %d", i, len(span.Children[i].Children))
		}
		for j := 0; j < children; j++ {
			if span.Children[i].Children[j].Name != fmt.Sprintf("t%d.%d", i, j) {
				t.Errorf("unexpected t%d.%d name: %v", i, j, span.Children[i].Children[j].Name)
			}
			if span.Children[i].Children[j].Duration <= 0 {
				t.Errorf("unexpected t%d.%d duration: %v", i, j, span.Children[i].Children[j].Duration)
			}
			if len(span.Children[i].Children[j].Children) != 0 {
				t.Errorf("unexpected t%d.%d Children length: %d", i, len(span.Children[i].Children[j].Children))
			}
		}
	}
}

func testRunSpan(ctx context.Context, name string, children ...string) {
	span := StartSpan(ctx, name)
	defer span.Finish()
	if len(children) == 0 {
		return
	}
	ctx = context.WithValue(ctx, SpanKey(), span)
	for _, child := range children {
		func() {
			span := StartSpan(ctx, child)
			defer span.Finish()
		}()
	}
}
