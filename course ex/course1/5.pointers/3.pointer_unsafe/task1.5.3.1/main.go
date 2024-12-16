package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(uint8ToInt8(128))
}

func uint8ToInt8(num uint8) int8 {
	return *(*int8)(unsafe.Pointer(&num))
}
