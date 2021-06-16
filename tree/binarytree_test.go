package tree

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8}
	res := BuildTree(arr)
	MidOrder(res)
}

func TestPostOrder(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8}
	res := BuildTree(arr)
	PostOrder(res)
}