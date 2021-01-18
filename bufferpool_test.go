package bufferpool_test

import (
	"bytes"
	"testing"

	"github.com/lestrrat-go/bufferpool"
	"github.com/stretchr/testify/assert"
)

func TestUsage(t *testing.T) {
	pool := bufferpool.New()

	buf := pool.Get()
	for i := 0; i < 10; i++ {
		if !assert.Equal(t, &bytes.Buffer{}, buf, `should be an empty buffer`) {
			return
		}
		pool.Release(buf)
	}
}
