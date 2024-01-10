package mpool

// DefaultMemPool .
var DefaultMemPool = NewSafeMpool(1024, 1024*1024*1024)

func Init(bufSize, freeSize int) {
	DefaultMemPool = NewSafeMpool(bufSize, freeSize)
}

// Malloc exports default package method.
func Malloc(size int) []byte {
	return DefaultMemPool.Malloc(size)
}

// Realloc exports default package method.
func Realloc(buf []byte, size int) []byte {
	return DefaultMemPool.Realloc(buf, size)
}

// Free exports default package method.
func Free(buf []byte) {
	DefaultMemPool.Free(buf)
}
