package queue

import "fmt"

type Node struct {
	Value		interface{}
	Next 	*Node
}

var(
	Size = 0
	queue = &Node{nil,nil}
)
func(n *Node) Push(value interface{}){
	if queue == nil {
		queue = &Node{value, nil}
	}else{
		n.Value = value
		n.Next = queue
	}
	Size++
}

func(n *Node) Pop() interface{} {
	if Size == 0 {
		return 0
	}
	if Size == 1{
		queue = nil
		Size--
		return n.Value
	}
	temp := n
	for (n.Next) != nil{
		temp = n
		n = n.Next
	}
	Size --
	return temp.Next.Value
}

//遍历节点
func (n *Node)Traverse() {
	if Size == 0 {
		fmt.Println("empty Queue!")
	}
	for n != nil {
		fmt.Printf("%v-> ",n.Value)
		n = n.Next
	}
}

/**
方法2
 */
type ListNode struct {
	Val		interface{}
	Next 	*ListNode
}

type ListQueue struct {
	head 	*ListNode
	tail 	*ListNode
	length 	int
}

func NewQueue() *ListQueue {
	return new(ListQueue)
}

func (l *ListQueue)Push(value interface{}) {
	node := &ListNode{value, nil}
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}
	l.length ++
}

func (l *ListQueue)Pop() (value interface{}) {
	if l.head == nil {
		return
	}
	value = l.tail.Val
	l.head = l.head.Next
	l.length --
	return
}

func(l *ListQueue) Traverse() {
	if l.head == nil {
		fmt.Printf("empty Queue!")
	}else{
		for cur := l.head; cur != nil; cur = cur.Next{
			fmt.Printf(" <-%+v",cur.Val)
		}

	}

}