package tracer

import (
	"context"
	"fmt"
	"time"
)

func ExampleContext() {
	ctx := context.Background()
	fmt.Println(ctx.Done())

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	fmt.Println(ctx.Done() != nil)

	k := struct{}{}
	ctx = context.WithValue(ctx, k, 333)
	fmt.Println(ctx.Done() != nil, ctx.Value(k))
	// Output:
	// <nil>
	// true
	// true 333
}

func ExampleStart() {
	ctx := Start(context.Background(), "name")
	Start(ctx, "child1")
	Start(ctx, "child2")
	Finish(ctx)

	tracer := Get(ctx)
	fmt.Println(
		tracer.Name, !tracer.At.IsZero(), tracer.Duration > 0, tracer.Tags, tracer.Logs,
		len(tracer.Children), tracer.Children[0].Name, tracer.Children[1].Name,
	)

	// Output:
	// name true true map[] [] 2 child1 child2
}

func ExampleTag() {
	ctx := Start(context.Background(), "name")

	Tag(ctx, "k", "v")
	fmt.Println(Get(ctx).Tags)

	DebugTag(ctx, "debugKey", "debugValue")
	fmt.Println(Get(ctx).Tags)
	ctx = SetDebug(ctx)
	DebugTag(ctx, "debugKey", "debugValue")
	fmt.Println(Get(ctx).Tags)

	// Output:
	// map[k:v]
	// map[k:v]
	// map[debugKey:debugValue k:v]
}

func ExampleTagString() {
	ctx := Start(context.Background(), "name")

	TagString(ctx, "k", map[string]string{"a": "b"})
	fmt.Println(Get(ctx).Tags)

	DebugTag(ctx, "debugKey", "debugValue")
	fmt.Println(Get(ctx).Tags)
	ctx = SetDebug(ctx)
	DebugTag(ctx, "debugKey", "debugValue")
	fmt.Println(Get(ctx).Tags)

	// Output:
	// map[k:{"a":"b"}]
	// map[k:{"a":"b"}]
	// map[debugKey:debugValue k:{"a":"b"}]
}

func ExampleLog() {
	ctx := Start(context.Background(), "name")

	Log(ctx, "a ", "b")
	fmt.Printf("%#v\n", Get(ctx).Logs)
	Logf(ctx, "a %s", "b")
	fmt.Printf("%#v\n", Get(ctx).Logs)

	DebugLog(ctx, "a ", "b")
	DebugLogf(ctx, "a %s", "b")
	fmt.Printf("%#v\n", Get(ctx).Logs)
	ctx = SetDebug(ctx)
	DebugLog(ctx, "a ", "b")
	DebugLogf(ctx, "a %s", "b")
	fmt.Printf("%#v\n", Get(ctx).Logs)

	// Output:
	// []string{"a b"}
	// []string{"a b", "a b"}
	// []string{"a b", "a b"}
	// []string{"a b", "a b", "a b", "a b"}
}
