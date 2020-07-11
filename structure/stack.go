package structure

type SingleNode struct {
	Data interface{}
	Next *SingleNode
}

type Stack struct {
	Length int
	Head *SingleNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(data interface{}) {
	s.Length++
	node:=&SingleNode{
		Data: data,
	}
	if s.Head == nil {
		s.Head=node
		return
	}
	node.Next,s.Head=s.Head,node
}

func (s *Stack) Pop() interface{} {
	s.Length--
	data:=s.Head.Data
	s.Head=s.Head.Next
	return data
}

func (s *Stack) Size() int {
	return s.Length
}
