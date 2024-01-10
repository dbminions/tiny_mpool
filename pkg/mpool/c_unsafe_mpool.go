package mpool

type UnsafeMemPool struct {
}

func (mp *UnsafeMemPool) Malloc(size int) []byte {
	//TODO implement me
	panic("implement me")
}

func (mp *UnsafeMemPool) Realloc(buf []byte, size int) []byte {
	//TODO implement me
	panic("implement me")
}

func (mp *UnsafeMemPool) Free(buf []byte) {
	//TODO implement me
	panic("implement me")
}
