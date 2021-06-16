package tree

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	n := &Node{Value: 5}
	n.Left = &Node{Value: 3}
	n.Right = &Node{Value: 6}
	n.Left.Left = &Node{Value: 2}
	n.Left.Right = &Node{Value: 4}
	n.Right.Left = &Node{Value:1}
	n.Print()
}

func TestPostOrder(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8}
	res := BuildTree(0, arr)
	PostOrder(res)
}