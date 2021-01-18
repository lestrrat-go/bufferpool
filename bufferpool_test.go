package bufferpool_test

import (
	"testing"

	"github.com/lestrrat-go/bufferpool"
	"github.com/stretchr/testify/assert"
)

func TestReuse(t *testing.T) {
	pool := bufferpool.New()

	buf1 := pool.Get()
	pool.Release(buf1)

	buf2 := pool.Get()
	if !assert.True(t, buf1 == buf2, `buffers should be the same`) {
		return
	}
}

func TestAlloc(t *testing.T) {
	pool := bufferpool.New()

	buf1 := pool.Get()
	buf2 := pool.Get()
	if !assert.False(t, buf1 == buf2, `buffers should be different`) {
		return
	}
}
