package question1

import (
	"GoLearn/testTool"
	"fmt"
	"reflect"
	"testing"
)

// 1. panic 在 defer 函数后先入后出运行后, 执行
func TestPanicAfterDefer(t *testing.T) {
	// 打印后
	// 打印中
	// 打印前
	// panic: 触发异常
	fmt.Println(deferCall())
}
func deferCall() int {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	//panic("触发异常")
	return 1
}

// 2 for range  循环副本创建
func TestForRangeCopy(t *testing.T) {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	// 循环的时候会创建每个元素的副本，
	// ⽽不是每个元素的引⽤，所以 m[key] = &val 取的 都是变量val的地址 ，
	// 所以最后 map 中的所有元素的值都是变量 val 的地址，
	// 因为最后 val 被赋值为 3，所有输出的都是3。
	for key, value := range slice {
		m[key] = &value
	}

	for key, val := range m {
		// 0 -> 3
		// 1 -> 3
		// 2 -> 3
		fmt.Println(key, "->", val)
	}
}

// 3. new 和 make 区别
func TestNewAndMake(t *testing.T) {

	// new(T) 和 make(T, args) 是Go语⾔内建函数，⽤来分配内存，但适⽤的类型不⽤。
	// new(T) 会为了 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的 值。
	// 换句话说就是，返回⼀个指针，该指针指向新分配的、类型为 T 的零值。适⽤于值类型，如 数组 、 结构体 等。
	// make(T, args) 返回初始化之后的T类型的值，也不是指针 *T ，是经过初始化之后的T的引⽤。
	// make() 只适⽤于 slice 、 map 和 channel 。
	i := new([]int)
	fmt.Println(i)
	makeT := make([]int, 0)
	makeT = append(makeT, 2)
	fmt.Println(makeT)
}

// 4 结构体比较
func TestStructCompare(t *testing.T) {

	// 1. 结构体只能⽐较是否相等，但是不能⽐较⼤⼩；
	// 2. 想同类型的结构体才能进⾏⽐较，结构体是否相同不但与属性类型有关，还与属性顺序相关；
	// 3. 如果struct的所有成员都可以⽐较，则该struct就可以通过==或!=进⾏⽐较是否相同，
	// ⽐较时逐个项 进⾏⽐较，如果每⼀项都相等，则两个结构体才相等，否则不相等；
	// 那有什么是可以⽐较的呢？
	//
	// 常⻅的有bool、数值型、字符、指针、数组等
	//
	// 不能⽐较的有
	//
	//slice、map、函数

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "11"}
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int

		m map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	fmt.Println(sm2, sm1)
	//if sm1 != sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
}

// 4 nil 切片 空切片集合

func TestNilSliceAndEmptySlice(t *testing.T) {
	var s []int

	type s4 []int
	s1 := s4{1, 1}
	s2 := reflect.TypeOf(s1).String()
	testTool.EqualTest(t, s, nil)
	testTool.EqualTest(t, s2, "question1.s4")
	testTool.EqualTest(t, reflect.TypeOf(s1).Kind().String(), "slice")
}

// 5 截取切片
func TestCutSlice(t *testing.T) {
	s := [3]int{1, 2, 3}

	a := s[:0]

	b := s[:2]

	fmt.Println(s[2:3])

	// [i:j:k]
	//  low <= high (length) <= max (cap)
	// 截取操作有带 2 个或者 3 个参数，形如：[i:j] 和 [i:j:k]，假设截取对象的底层数组⻓度为 j。
	// 在操作符 [i:j] 中，如果 i 省略，默认 0，如果 j 省略，默认底层数组的⻓度.
	// 操作符 [i:j:k]，k 主要是⽤来限制切⽚的容量，但是不能⼤于数组的⻓度 l，
	//  截取得到的切⽚⻓度 和容量计算⽅法是 j-i、k-i。
	c := s[0:1:3]
	fmt.Println("cap c", cap(c), "cap s", cap(s))
	fmt.Println("a-", a, "b-", b, "c-", c)
}

// 6 变量命名规则
func TestVarBatch(t *testing.T) {
	var x int

	f := func() (int, int) {
		return 2, 1
	}

	//x, _ := f() 错

	x, _ = f() // 对
	//
	//x, y := f() // 对
	fmt.Println(x)
	//fmt.Println(y)

	//x, y = f() // 错
}

// 7  接口类型断言
func TestAssertType(t *testing.T) {
	type A interface{ ShowA() int }
	type B interface {
		ShowB() int
	}

	var a A = Work{3}

	s := a.(Work)
	fmt.Println(s.ShowA())
	fmt.Println(s.ShowB())
}

type Work struct{ i int }

func (w Work) ShowA() int { return w.i + 10 }

func (w Work) ShowB() int { return w.i + 20 }
