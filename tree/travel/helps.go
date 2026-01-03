package node

//中序遍历树

func (this *TreeNode) Travels() {
	if this == nil {
		return
	}
	this.Left.Travels()
	this.Print()
	this.Right.Travels()
}
