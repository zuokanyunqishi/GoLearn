package str

import (
	"fmt"
	"testing"
	"unsafe"
)

// try to modify string fail
func TestModifyStr(t *testing.T) {
	var s = "hello"
	fmt.Println("origin str:", s)
	modifystring(&s)
	fmt.Println(s)
}

func modifystring(s *string) {
	//取出第一个8宇节的值
	p := (*uintptr)(unsafe.Pointer(s))
	//获取底层数组的地址
	var array = (*[5]byte)(unsafe.Pointer(*p))
	var l = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Sizeof((*uintptr)(nil))))
	for i := 0; i < (*l); i++ {
		fmt.Printf("%p => %c\n", &((*array)[i]), (*array)[i])
		p1 := &((*array)[i])
		v := (*p1)
		(*p1) = v + 1 // try to change the character
	}
}
