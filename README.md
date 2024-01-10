## Tiny Memory Pool

Memory pool is a memory management technique that pre-allocates a large amount of memory from the system and then
provides it to the client when needed. It is also known as **fixed-size block allocator** or **region-based memory
allocator**. It is a special case of **object pool pattern**. It is used to manage memory in garbage collected
languages where the garbage collector does not provide efficient memory management.

### Types of Memory Pool

- [Safe Memory Pool](./pkg/mpool/b_safe_mpool.go): Here we are using Go primitives like `make([]byte, size)` and
  `sync.Pool` to allocate and deallocate memory.
- [Unsafe Memory Pool](./pkg/mpool/c_unsafe_mpool.go): Here we are using `unsafe` package to allocate and deallocate
  memory. 