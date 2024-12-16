package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println(binaryStringToFloat("00111110001000000000000000000000"))
}

func binaryStringToFloat(binary string) float32 {
	var number uint32
	num1, err := strconv.ParseUint(binary,2,0)
	if err != nil {
		panic(err)
	}
	number = uint32(num1)
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}
