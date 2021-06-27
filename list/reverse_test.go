package list

import "testing"

func TestReverse(t *testing.T) {
	array := []int{1,2,3,4,5}
	l := ConvertListNode(array)
	arr := ReverseList(l)
	arr.PrintListNode()
}