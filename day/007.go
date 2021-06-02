package day

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