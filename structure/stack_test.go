package structure

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
		fmt.Println(stack.Size())
	}

	for i := 0; i < 10; i++ {
		fmt.Println(stack.Pop(), stack.Size())
	}

}
