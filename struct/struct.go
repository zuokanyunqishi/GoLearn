package main

import (
	"encoding/json"
	"fmt"
)

// 结构体
type treeNode struct {
	value       int
	left, right *treeNode
}

// 设置value
func (this *treeNode) setValue(value int) {
	this.value = value
}

// 归属结构体的方法
func (this *treeNode) print() {

	fmt.Print(this.value)

}

//中序遍历树

func (this *treeNode) travels() {
	if this == nil {
		return
	}
	this.left.travels()
	this.print()
	this.right.travels()
}

// 创建树
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

type cInt int

func (c *cInt) change() {
	*c = *c + 1

}

func (c *cInt) String() string {
	return fmt.Sprintf("%v 你FDFJD", *c+6)
}

type test struct {
	Name *string
	Age  int
}

func (t test) print1() {
	//
	t.Age = 15
	fmt.Printf("test print 1 ,t address %p \n", &t)
	fmt.Println("test print 1 name ", *t.Name)
	fmt.Println("test print 1 age ", t.Age)
}

func (t *test) print2() {
	print2str := "abc"

	t.Name = &print2str
	fmt.Printf("*t 值的 地址是 %p \n", &(*t))
	t.Age = 13
	fmt.Println("test print 2 ", *(t.Name))
	fmt.Println("test print 2 value ", t.Name)
	fmt.Println("test print 2 ", &t.Age)

}
func main() {
	var root treeNode

	root = treeNode{value: 4}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = createNode(4)
	root.right.right = new(treeNode)

	//结构体切片
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	root.setValue(15)
	root.travels()
	for _, v := range nodes {
		v.travels()
		fmt.Print("\n")
	}

	newStruct()

	var int2 cInt
	// 实现 String() 方法 ，自动调用
	fmt.Println(&int2)

	var test1 test
	strName := "测试名字"
	test1.Name = &strName
	test1.Age = 10

	fmt.Printf("test 1变量的地址 %p \n", &test1)
	// 值接收者是 复制 一份新的 到函数栈空间
	test1.print1()
	fmt.Println()
	fmt.Println("age", test1.Age)
	test1.print2()

	fmt.Println("test1 age 变量 地址 ，", &test1.Age)

}

func newStruct() {
	type student struct {
		Age     int               `json:"age"`
		Name    string            `json:"name"`
		Score   []int             `json:"score"`
		Subject map[string]string `json:"subject"`
	}

	// new 返回指针类型
	student1 := new(student)
	student1.Age = 3
	student1.Name = "abc"
	student1.Score = []int{3, 45, 56}
	student1.Subject = map[string]string{"A": "语文"}
	fmt.Println(student1)

	// var 实例化结构体
	var student2 = student{}
	student2.Age = 5
	student2.Name = "小明"
	student2.Score = make([]int, 4)
	student2.Subject = make(map[string]string)
	fmt.Println(student2)

	student3 := student2
	// 结构体是值类型
	fmt.Printf("studnet2,内存地址是 %p \n sudent3,内存地址是 %p \n", &student2, &student3)

	// student4 是个指针变量
	student4 := &student{}
	(*student4).Subject = map[string]string{"语文": "我是", "数学": "344"}
	(*student4).Age = 45
	(*student4).Name = "老王"

	// 结构体在内存是连续分配的

	fmt.Printf("Age *P %p ,\nName *p %p ,\nScore *p %p", &(*student4).Age, &(*student4).Name, &(*student4).Score)

	// struct tag 序列化
	data, _ := json.Marshal(*student4)
	fmt.Println(string(data))

}
