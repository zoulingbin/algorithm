package list

import "math"

//跳表
const LEVEL = 10 //最高层数

//跳表节点结构体
type skipListNode struct {
	data 	interface{}
	score 	int
	//层数
	level 	int
	//每层前进指针
	forwards []*skipListNode
}

type SkipList struct {
	//跳表头结点
	head 	*skipListNode
	//跳表当前层数
	level 	int
	//跳表长度
	length 	int
}

func NewSkipListNode(data interface{}, score int, level int) *skipListNode {
	//头节点
	return &skipListNode{data: data, score: score, level: level, forwards: make([]*skipListNode, level, level)}
}

func NewSkipList() *SkipList {
	head := NewSkipListNode(0, math.MinInt32, LEVEL)
	return &SkipList{head: head, level: 1, length: 0}
}