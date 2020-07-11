package structure_test

import (
	"fmt"
	"testing"

	"go-algorithm/structure"
)

func TestNewList(t *testing.T) {
	list := structure.NewList()
	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	fmt.Println(list.Insert(1+1i,1))
	fmt.Println(list.Insert(2+1i,1))
	fmt.Println(list.Insert(1+2i,2))
	fmt.Println(list.Insert(13,14))
	fmt.Println(list.Insert(0,0))
	fmt.Println(list.Insert(1000,1000))

	fmt.Println(list.Select(5).Data)
	fmt.Println(list.Delete(5))

	fmt.Println(list)
}
