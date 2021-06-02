package list

import "fmt"

type ListNode struct {
	Val 	int
	Next 	*ListNode
}

func Reverse(head *ListNode) *ListNode{
	cur := head
	var pre *ListNode

	for cur != nil{
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func ReverseList (head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

//构建链表
func ConvertListNode(s []int) *ListNode {
	sLen := len(s)
	if sLen == 0 {
		return nil
	}
	listNode := &ListNode{
		Val: s[0],
	}
	temp := listNode
	for i := 1; i < sLen; i++ {
		temp.Next = &ListNode{
			Val: s[i],
		}
		temp = temp.Next
	}
	return listNode
}

//打印链表
func (l *ListNode) PrintListNode() {
	var s []int
	pointer := l
	for pointer != nil {
		s = append(s, pointer.Val)
		pointer = pointer.Next
	}
	for _, v := range s {
		fmt.Println(v)
	}
}