package node

import "fmt"

// TreeNode 结构体
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// SetValue 设置value
func (this *TreeNode) SetValue(value int) {
	this.Value = value
}

// Print 归属结构体的方法
func (this *TreeNode) Print() {

	fmt.Print(this.Value)

}

// CreateNode 创建树
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}
