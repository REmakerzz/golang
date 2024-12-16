package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool
	var n int
	var n8 int8
	var n16 int16
	var n32 int32
	var n64 int64
	var nuint uint
	var nuint8 uint8

	fmt.Println("bool size:", sizeOfBool(b))
	fmt.Println("int size:", sizeOfInt(n))
	fmt.Println("int8 size:", sizeOfInt8(n8))
	fmt.Println("int16 size:", sizeOfInt16(n16))
	fmt.Println("int32 size:", sizeOfInt32(n32))
	fmt.Println("int64 size:", sizeOfInt64(n64))
	fmt.Println("uint size:", sizeOfUint(nuint))
	fmt.Println("uint8 size:", sizeOfUint8(nuint8))
}

func sizeOfBool(b bool) int {
	result := unsafe.Sizeof(b)
	return int(result)
}

func sizeOfInt(n int) int {
	result := unsafe.Sizeof(n)
	return int(result)
}

func sizeOfInt8(n int8) int {
	result := unsafe.Sizeof(n)
	return int(result)
}

func sizeOfInt16(n int16) int {
	result := unsafe.Sizeof(n)
	return int(result)
}

func sizeOfInt32(n int32) int {
	result := unsafe.Sizeof(n)
	return int(result)
}

func sizeOfInt64(n int64) int {
	result := unsafe.Sizeof(n)
	return int(result)
}

func sizeOfUint(n uint) int {
	result := unsafe.Sizeof(n)
	return int(result)
}

func sizeOfUint8(n uint8) int {
	result := unsafe.Sizeof(n)
	return int(result)
}
