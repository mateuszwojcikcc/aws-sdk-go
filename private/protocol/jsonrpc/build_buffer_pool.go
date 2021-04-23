package jsonrpc

import (
	"bytes"
	"sync"
)

var buildBufferPool = bufferPool{}

type bufferPool struct {
	pool sync.Pool
}

// get returns AggregatedRecord from pool.
// The AggregatedRecord must be put to pool after use.
func (pp *bufferPool) get() *bytes.Buffer {
	v := pp.pool.Get()
	if v == nil {
		return &bytes.Buffer{}
	}

	return v.(*bytes.Buffer)
}

// put returns p to pool.
func (pp *bufferPool) put(p *bytes.Buffer) {
	p.Reset()

	pp.pool.Put(p)
}
