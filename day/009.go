package day

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	cur, pre := head, head
	for cur.Next != nil {
		if cur.Value == val {
			pre.Next = cur.Next
		}else{
			pre = cur
		}
		cur = cur.Next
	}
	return head
}
