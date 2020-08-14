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
	if s.Length <= 0 {
		return nil
	}
	s.Length--
	defer func() {
		s.Head=s.Head.Next
	}()
	return s.Head.Data
}

func (s *Stack) Size() int {
	return s.Length
}
