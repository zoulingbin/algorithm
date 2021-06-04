package day

func MiddleNode(head *ListNode) *ListNode {
	fast, slow := head, head.Next
	for  {
		if fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}else if fast.Next == nil {
			return slow.Next
		}else {
			return slow
		}
	}
}