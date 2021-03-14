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
func(n *Node) Push(value interface{}) bool{
	if queue == nil {
		queue = &Node{value, nil}
		Size ++
		return true
	}
	n.Value = value
	n.Next = queue
	Size++
	return true
}

func(n *Node) Pop() (interface{}, bool){
	if Size == 0 {
		return 0,false
	}
	if Size == 1{
		queue = nil
		Size--
		return n.Value,true
	}
	temp := n
	for (n.Next) != nil{
		temp = n
		n = n.Next
	}
	Size --
	return temp.Next.Value, true
}

//遍历节点
func (n *Node)traverse() {
	if Size == 0 {
		fmt.Println("empty Queue!")
	}
	for n != nil {
		fmt.Printf("%v-> ",n.Value)
		n = n.Next
	}
}