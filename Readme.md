# tracer
a tracer for golang.

[![Build Status](https://travis-ci.org/lovego/tracer.svg?branch=master)](https://travis-ci.org/lovego/tracer)
[![Coverage Status](https://img.shields.io/coveralls/github/lovego/tracer/master.svg)](https://coveralls.io/github/lovego/tracer?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/lovego/tracer)](https://goreportcard.com/report/github.com/lovego/tracer)
[![GoDoc](https://godoc.org/github.com/lovego/tracer?status.svg)](https://godoc.org/github.com/lovego/tracer)

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

## Documentation
[https://godoc.org/github.com/lovego/tracer](https://godoc.org/github.com/lovego/tracer)
