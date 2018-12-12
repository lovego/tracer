package tracer

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestStartSpanReturnNil(t *testing.T) {
	if got := StartSpan(nil, "name"); got != nil {
		t.Errorf("unexpected: %v", got)
	}
	if got := StartSpan(context.Background(), "name"); got != nil {
		t.Errorf("unexpected: %v", got)
	}
}

func TestStartSpan(t *testing.T) {
	var span = &Span{Name: "root", At: time.Now()}
	func() {
		defer span.Finish()
		ctx := Context(context.Background(), span)

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
				t.Errorf("unexpected t%d.%d Children length: %d", i, j, len(span.Children[i].Children[j].Children))
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
	ctx = Context(ctx, span)
	for _, child := range children {
		func() {
			span := StartSpan(ctx, child)
			defer span.Finish()
		}()
	}
}

func TestSpanTag(t *testing.T) {
	var span *Span
	if span.GetDebug() {
		t.Errorf("unexpected Debug(): true")
	}
	span.Tag("k", "v")

	span = (&Span{}).Tag("k", "v")
	if len(span.Tags) != 1 || span.Tags["k"] != "v" {
		t.Errorf("unexpected Tags: %v", span.Tags)
	}
}

func TestSpanDebugTag(t *testing.T) {
	span := (&Span{}).DebugTag("k", "v")
	if len(span.Tags) != 0 {
		t.Errorf("unexpected Tags: %v", span.Tags)
	}
	span.SetDebug(true).DebugTag("k", "v")
	if len(span.Tags) != 1 || span.Tags["k"] != "v" {
		t.Errorf("unexpected Tags: %v", span.Tags)
	}
	if !span.GetDebug() {
		t.Errorf("unexpected Debug(): false")
	}
}

func TestSpanLog(t *testing.T) {
	var span *Span
	span.Log("a", "b")

	span = (&Span{}).Log("a ", "b")
	if len(span.Logs) != 1 || span.Logs[0] != "a b" {
		t.Errorf("unexpected Logs: %v", span.Logs)
	}
}

func TestSpanLogf(t *testing.T) {
	var span *Span
	span.Logf("a %s", "b")

	span = (&Span{}).Logf("a %s", "b")
	if len(span.Logs) != 1 || span.Logs[0] != "a b" {
		t.Errorf("unexpected Logs: %v", span.Logs)
	}
}

func TestSpanDebugLog(t *testing.T) {
	var span *Span
	span.DebugLog("a", "b")

	span = (&Span{}).SetDebug(true).DebugLog("a ", "b")
	if len(span.Logs) != 1 || span.Logs[0] != "a b" {
		t.Errorf("unexpected Logs: %v", span.Logs)
	}
}

func TestSpanDebugLogf(t *testing.T) {
	var span *Span
	span.DebugLogf("a %s", "b")

	span = (&Span{}).SetDebug(true).DebugLogf("a %s", "b")
	if len(span.Logs) != 1 || span.Logs[0] != "a b" {
		t.Errorf("unexpected Logs: %v", span.Logs)
	}
}
