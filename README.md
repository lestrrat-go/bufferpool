# bufferpool

Very simple bytes.Buffer pool using sync.Pool.

[![Build Status](https://travis-ci.org/lestrrat-go/bufferpool.svg?branch=master)](https://travis-ci.org/lestrrat-go/bufferpool)

[![GoDoc](https://godoc.org/github.com/lestrrat-go/bufferpool?status.svg)](https://godoc.org/github.com/lestrrat-go/bufferpool)


I got tired of writing the same sync.Pool for byte.Buffer objects.

# SYNOPSIS


```
import "github.com/lestrrat-go/bufferpool"

var pool = bufferpool.New()

func main() {
    buf := pool.Get()
    defer pool.Release(buf)

    // ...
}
```
