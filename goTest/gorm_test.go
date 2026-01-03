package goTest

import (
	node "GoLearn/tree/travel"
	"fmt"
	"log"
	"net/http"
	"testing"
)

// 逻辑测试
func TestFun(t *testing.T) {
	test := []struct{ a, b, c int }{
		{3, 4, 6},
		{356, 68, 9},
	}
	for _, parm := range test {
		result := add(parm.a, parm.b, parm.c)
		if result > 50 {
			t.Errorf(`超过50`)
		}
	}

}

// 性能测试 go test -bench . -cpuprofile cpu.out 生成性能宝宝
// go tool pprof cpu.out  分析cpu文件
// web 生成函数调用图
func BenchmarkTest(b *testing.B) {

	for i := 0; i < b.N; i++ {
		abc()
	}
}

// 测试http服务器
func TestHttp(t *testing.T) {
	http.HandleFunc(`/index`, response)
	log.Fatal(http.ListenAndServe(":8080", nil).Error())

}

func response(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`helloWord`))
}

func add(num, num1, num2 int) int {
	return num + num1 + num2
}

func abc() {
	var root node.TreeNode

	root = node.TreeNode{Value: 4}
	root.Left = &node.TreeNode{}
	root.Right = &node.TreeNode{5, nil, nil}
	root.Right.Left = node.CreateNode(4)
	root.Right.Right = new(node.TreeNode)

	//结构体切片
	nodes := []node.TreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	root.SetValue(15)
	root.Travels()
	for _, v := range nodes {
		v.Travels()
		fmt.Print("\n")
	}
}
