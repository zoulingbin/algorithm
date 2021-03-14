package list

type ListNode struct {
	Value 	interface{}
	Next 	*ListNode
}

func (l *ListNode) Reverse() *ListNode{
	cur := l
	var pre *ListNode

	for cur != nil{
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
