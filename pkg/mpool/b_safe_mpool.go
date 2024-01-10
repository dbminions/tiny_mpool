// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mpool

import (
	"sync"
)

type SafeMemPool struct {
	bufSize  int
	freeSize int
	pool     *sync.Pool
}

func NewSafeMpool(bufSize, freeSize int) Allocator {
	if bufSize <= 0 {
		bufSize = 64
	}
	if freeSize <= 0 {
		freeSize = 64 * 1024
	}
	if freeSize < bufSize {
		freeSize = bufSize
	}

	mp := &SafeMemPool{
		bufSize:  bufSize,
		freeSize: freeSize,
		pool:     &sync.Pool{},
	}
	mp.pool.New = func() interface{} {
		buf := make([]byte, bufSize)
		return &buf
	}

	return mp
}

func (mp *SafeMemPool) Malloc(size int) []byte {
	if size > mp.freeSize {
		return make([]byte, size)
	}
	pbuf := mp.pool.Get().(*[]byte)
	n := cap(*pbuf)
	if n < size {
		*pbuf = append((*pbuf)[:n], make([]byte, size-n)...)
	}
	return (*pbuf)[:size]
}

func (mp *SafeMemPool) Realloc(buf []byte, size int) []byte {
	if size <= cap(buf) {
		return buf[:size]
	}

	if cap(buf) < mp.freeSize {
		pbuf := mp.pool.Get().(*[]byte)
		n := cap(buf)
		if n < size {
			*pbuf = append((*pbuf)[:n], make([]byte, size-n)...)
		}
		*pbuf = (*pbuf)[:size]
		copy(*pbuf, buf)
		mp.Free(buf)
		return *pbuf
	}
	return append(buf[:cap(buf)], make([]byte, size-cap(buf))...)[:size]
}

func (mp *SafeMemPool) Free(buf []byte) {
	if cap(buf) > mp.freeSize {
		return
	}
	mp.pool.Put(&buf)
}
