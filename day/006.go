package day

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil{
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	tem := &ListNode{}
	if l1.Value <= l2.Value {
		tem = l1
		tem.Next = MergeTwoLists(l1.Next, l2)
	} else {
		tem = l2
		tem.Next = MergeTwoLists(l1, l2.Next)
	}
	return tem
}
