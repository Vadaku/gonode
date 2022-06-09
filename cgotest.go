package main

/*
#include <stdio.h>
#include <stdlib.h>

char* getGPU(char* s);
#cgo LDFLAGS: -L./lib -lgetgpu -Wl,-rpath=./lib

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func Example() {
	pcontent := C.CString("test")
	testhash := C.getGPU(pcontent)
	C.free(unsafe.Pointer(pcontent))

	fmt.Printf("%s\n", C.GoString(testhash))
}
