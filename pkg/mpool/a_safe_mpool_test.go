package mpool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSafeMpool(t *testing.T) {
	mp := NewSafeMpool(1024, 1024*1024*1024)
	buf := mp.Malloc(1024)
	buf[0] = 'h'
	buf[1] = 'e'
	buf[2] = 'l'
	buf[3] = 'l'
	buf[4] = 'o'

	assert.Equal(t, 1024, len(buf))
	mp.Free(buf)
}
