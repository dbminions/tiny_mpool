package mpool

import "testing"

func TestNewSafeMpool(t *testing.T) {
	mp := NewSafeMpool(1024, 1024*1024*1024)
	buf := mp.Malloc(1024)
	mp.Free(buf)
}
