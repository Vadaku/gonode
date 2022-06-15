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
	"fmt"
	"time"
	"unsafe"
)

//C Function Wrapper
func CudaHash(toHash string) string {
	pcontent := C.CString(toHash)
	testhash := C.cudaHash(pcontent)
	C.free(unsafe.Pointer(pcontent))

	return C.GoString(testhash)
}

func HelloCpp(source string, data string, target string) {
	source1 := C.CString(source)
	target1 := C.CString(target)
	data1 := C.CString(data)
	user := C.CString("anon")

	start := time.Now()
	C.cudaMine(source1, target1, 0, data1, user, 1000, 5000)
	elapsed := time.Since(start)
	fmt.Printf("\033[32mTime taken %s\033[0m\n", elapsed)

	C.free(unsafe.Pointer(source1))
	C.free(unsafe.Pointer(target1))
	C.free(unsafe.Pointer(data1))
	C.free(unsafe.Pointer(user))

}
