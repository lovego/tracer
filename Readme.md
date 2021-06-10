# tracer
a tracer for golang.

[![Build Status](https://github.com/lovego/tracer/actions/workflows/go.yml/badge.svg)](https://github.com/lovego/tracer/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/lovego/tracer/badge.svg?branch=master)](https://coveralls.io/github/lovego/tracer)
[![Go Report Card](https://goreportcard.com/badge/github.com/lovego/tracer)](https://goreportcard.com/report/github.com/lovego/tracer)
[![Documentation](https://pkg.go.dev/badge/github.com/lovego/tracer)](https://pkg.go.dev/github.com/lovego/tracer@v0.0.1)

## Install
`$ go get github.com/lovego/tracer`

## Usage
```go
func main() {
  ctx := tracer.Start(context.Background(), "main")
  defer tracer.Finish(ctx)

  tracer.Tag(ctx, "key", "value")
  tracer.Logf(ctx, "%v %v", time.Now(), "event")
 
  fmt.Println(tracer.Get(ctx))
}
```

