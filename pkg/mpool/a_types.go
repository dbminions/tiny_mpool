package mpool

type Allocator interface {
	Malloc(size int) []byte
	Realloc(buf []byte, size int) []byte
	Free(buf []byte)
}

var _ Allocator = (*SafeMemPool)(nil)
var _ Allocator = (*UnsafeMemPool)(nil)
