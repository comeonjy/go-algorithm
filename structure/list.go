package structure

import (
	"fmt"
)

type Node struct {
	Data        interface{}
	Left, Right *Node
}

type List struct {
	Length     int
	Head, Tail *Node
}

func NewList() *List {
	return &List{}
}

// 链表追加
func (l *List) Append(data interface{}) {
	l.Length++
	node := &Node{
		Data: data,
	}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Left = l.Tail
		l.Tail.Right = node
		l.Tail = node
	}

}

// 链表插入
func (l *List) Insert(data interface{}, n int) bool {
	if n-1 > l.Length || n <= 0 {
		return false
	}
	l.Length++
	node := &Node{
		Right: l.Head,
		Data:  data,
	}
	if n == 1 {
		l.Head.Left = node
		l.Head = node
		return true
	}

	temp := l.Head
	for i := 0; i < n-2; i++ {
		temp = temp.Right
	}

	if temp.Right != nil {
		node.Right = temp.Right
		temp.Right.Left = node
	}
	node.Left = temp
	temp.Right = node

	return true
}

// 按值查找
func (l *List) Select(value interface{}) *Node {
	temp:=l.Head
	for i := 0; i < l.Length; i++ {
		if temp.Data==value {
			return temp
		}
		temp=temp.Right
	}
	return nil
}

// 按值删除
func (l *List) Delete(value interface{}) bool {
	temp:=l.Head
	for i := 0; i < l.Length; i++ {
		if temp.Data == value {
			if temp.Right!=nil {
				temp.Right.Left=temp.Left
				temp.Left.Right=temp.Right
			}else{
				temp.Left.Right=nil
			}
			l.Length--
			return true
		}
		temp=temp.Right
	}
	return false
}

func (l *List) String() string {
	var str string
	temp := l.Head
	for i := 0; i < l.Length; i++ {
		str += fmt.Sprintf("%v<-->", temp.Data)
		temp = temp.Right
	}
	return str
}
