package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

char* cudaHash(char* s);
void hello(char* msg);
    void cudaMine(char* s, char* t, uint64_t nonce,
    char* d, char* u, int numMessagesPerIteration, uint64_t timestamp);
#cgo LDFLAGS: -L./lib -lcudahash -Wl,-rpath=./lib
#cgo CXXFLAGS: -std=c++14 -I.
*/
import "C"

import (
	"unsafe"
)

//C Function Wrapper
func CudaHash(toHash string) string {
	pcontent := C.CString(toHash)
	testhash := C.cudaHash(pcontent)
	C.free(unsafe.Pointer(pcontent))

	return C.GoString(testhash)
}

func HelloCpp() {
	source := C.CString("hello")
	target := C.CString("21e8")
	data := C.CString("data")
	user := C.CString("anon")

	C.cudaMine(source, target, 0, data, user, 1000, 5000)
	C.free(unsafe.Pointer(source))
	C.free(unsafe.Pointer(target))
	C.free(unsafe.Pointer(data))
	C.free(unsafe.Pointer(user))

}
