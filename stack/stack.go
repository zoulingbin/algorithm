package stack

import "fmt"

type Node struct {
	Value 	interface{}
	Next	*Node
}

//栈顶节点
type LinkedListStack struct {
	topNode *Node
}

func NewLinkListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

//func NewNode() *Node{
//	return new(Node)
//}

func (l *LinkedListStack) Push(v interface{}){
	l.topNode = &Node{v,l.topNode}
}

func (l *LinkedListStack) Pop() interface{}{
	if l.topNode == nil {
		return nil
	}
	v := l.topNode.Value.(*Node).Value
	l.topNode = l.topNode.Next
	return v
}

func (l *LinkedListStack) traverse() {
	if l.topNode == nil {
		fmt.Println("empty stack")
	}else{
		cur := l.topNode.Next
		for cur != nil {
			fmt.Printf("%d -> ",cur.Value)
			cur = cur.Next
		}
	}

}


