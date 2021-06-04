package day

import "fmt"

type ListNode struct {
	Value 	int
	Next 	*ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Value == cur.Next.Value {
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}
	return head
}

func (l *ListNode) printListNode() {
	var s []int
	pointer := l
	for pointer != nil {
		s = append(s, pointer.Value)
		pointer = pointer.Next
	}
	for _, v := range s {
		fmt.Println(v)
	}
}

// 转换链表
func convertListNode(s []int) *ListNode {
	sLen := len(s)
	if sLen == 0 {
		return nil
	}
	listNode := &ListNode{
		Value: s[0],
	}
	temp := listNode
	for i := 1; i < sLen; i++ {
		temp.Next = &ListNode{
			Value: s[i],
		}
		temp = temp.Next
	}
	return listNode
}