package algorithm

import (
	"container/list"
	"errors"
	"sync"
)

//最近未使用算法
type Lru struct {
	Max 	int
	l 		*list.List
	Call	func(key interface{}, value interface{})
	cache map[interface{}]*list.Element
}

type Node struct {
	Key,Value 	interface{}
}

var mu *sync.Mutex

func NewLru(n int) *Lru{
	return &Lru{
		Max: n,
		l:	list.New(),
		cache: make(map[interface{}]*list.Element),
	}
}

func (l *Lru) Set(key interface{}, value interface{}) error {
	if l.l == nil {
		return errors.New("LRU未初始化")
	}
	mu.Lock()
	defer mu.Unlock()
	if element, ok := l.cache[key]; ok {
		element.Value.(*Node).Value = value
		l.l.MoveToFront(element)
		return nil
	}
	ele := l.l.PushFront(&Node{
		Key: key,
		Value: value,
	})
	l.cache[key] = ele
	//超出长度 删除最后的元素
	if l.Max != 0 && l.l.Len() > l.Max {
		if e := l.l.Back(); e != nil{
			l.l.Remove(e)
			node := e.Value.(*Node)
			delete(l.cache,node.Key)
			if l.Call != nil{
				l.Call(node.Key, node.Value)
			}
		}
	}
	return nil
}

func (l *Lru)Get(key interface{}) (value interface{}, ok bool){
	if l.cache == nil {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if element, ok := l.cache[key]; ok {
		l.l.MoveToFront(element)
		return element.Value.(*Node).Value, true
	}
	return
}

func (l *Lru)Del(key interface{}) (ok bool){
	if l.cache == nil {
		return false
	}
	mu.Lock()
	defer mu.Unlock()
	if element, ok := l.cache[key]; ok{
		delete(l.cache, element)
		if e := l.l.Back(); e != nil{
			l.l.Remove(e)
			delete(l.cache, key)
			if l.Call != nil {
				node := e.Value.(*Node)
				l.Call(node.Key, node.Value)
				return true
			}
		}
	}
	return false
}