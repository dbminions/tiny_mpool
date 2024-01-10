package main

import (
	"tiny_mpool/pkg/mpool"
)

func main() {
	mp := mpool.NewSafeMpool(1024, 1024*1024*1024)
	buf := mp.Malloc(1024)
	buf[0] = 'h'
	buf[1] = 'e'
	buf[2] = 'l'
	buf[3] = 'l'
	buf[4] = 'o'

	mp.Free(buf)
}
