package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("%T %d", 89.000, unsafe.Sizeof(89.4444))

}
