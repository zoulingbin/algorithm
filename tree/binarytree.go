package tree

import "fmt"

type Node struct {
	Value 	int
	Left 	*Node
	Right 	*Node
}

//构建二叉树
var i = -1
func BuildTree(array []int) *Node {
	i += 1
	if i >= len(array) {
		return nil
	}
	t := new(Node)
	if array[i] != 0 {
		t.Value = array[i]
		t.Left = BuildTree(array)
		t.Right = BuildTree(array)
	} else {
		return nil
	}
	return t
}

//打印二叉树
func (tree *Node) Print() {
	if tree == nil {
		return
	}
	fmt.Println(tree.Value)
	tree.Left.Print()
	tree.Right.Print()
}

//前序遍历
func PreOrder(tree *Node)  {
	if tree == nil {
		return
	}
	fmt.Println(tree.Value)
	// 再打印左子树
	PreOrder(tree.Left)
	// 再打印右子树
	PreOrder(tree.Right)
}

// 中序遍历
func MidOrder(tree *Node){
	if tree ==nil{
		return
	}
	// 先打印左子树
	MidOrder(tree.Left)
	// 再打印根节点
	fmt.Print(tree.Value," ")
	// 再打印右字树
	MidOrder(tree.Right)
}

// 后序遍历
func PostOrder(tree *Node){
	if tree ==nil{
		return
	}
	// 先打印左子树
	PostOrder(tree.Left)
	// 再打印右字树
	PostOrder(tree.Right)
	// 再打印根节点
	fmt.Print(tree.Value," ")
}