package tree

type BinarySearchTree struct {
	Root 	*BinarySearchTreeNode
}

type BinarySearchTreeNode struct {
	Value 	int
	Left 	*BinarySearchTreeNode
	Right 	*BinarySearchTreeNode
}

//插入
func (tree *BinarySearchTree) Add(value int) {
	if tree.Root == nil {
		tree.Root = &BinarySearchTreeNode{Value: value}
		return
	}
	tree.Root.Add(value)
}

func (node *BinarySearchTreeNode) Add(value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &BinarySearchTreeNode{Value: value}
		}else{
			node.Left.Add(value)
		}
	}else if value > node.Value{
		if node.Right == nil {
			node.Right = &BinarySearchTreeNode{Value: value}
		}else{
			node.Right.Add(value)
		}
	}
}

//查找
func (tree *BinarySearchTree) Find (value int) *BinarySearchTreeNode {
	if tree.Root == nil {
		return nil
	}
	return tree.Root.Find(value)
}

func (node *BinarySearchTreeNode) Find(value int) *BinarySearchTreeNode {
	if value == node.Value {
		return node
	}
	if value < node.Value {
		if node.Left == nil {
			return node.Left
		}
		return node.Left.Find(value)
	}else {
		if node.Right == nil {
			return node.Right
		}
		return node.Right.Find(value)
	}
}

//删除