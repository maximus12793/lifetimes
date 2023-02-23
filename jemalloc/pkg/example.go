package lifetimes

import (
	"runtime/cgo"
	"unsafe"
)

/*
#cgo CFLAGS: -std=gnu99
#include <jemalloc/jemalloc.h>

// Define a custom allocator function that uses jemalloc to allocate memory.
void* jemalloc(size_t size) {
    return je_malloc(size);
}

// Define a custom free function that uses jemalloc to free memory.
void je_free(void* ptr) {
    je_free(ptr);
}
*/
import "C"

// Set jemalloc as the memory allocator for Go.
func init() {
	cgo.SetAllocator(C.jemalloc)
}

// Allocate memory for a new object using jemalloc.
func New(size uintptr) unsafe.Pointer {
	return C.jemalloc(C.size_t(size))
}

// Deallocate memory for an object using je_free.
func Delete(ptr unsafe.Pointer) {
	C.je_free(ptr)
}
