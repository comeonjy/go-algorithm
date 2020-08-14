package structure

import (
	"fmt"
)

type Tree struct {
	Data        interface{}
	Left, Right *Tree
}

func Add(data interface{})  {

}

// 递归先序遍历
func PreOrder(root *Tree)  {
	if root!=nil{
		fmt.Println(root.Data)
		PreOrder(root.Left)
		PreOrder(root.Right)
	}
}

// 栈中序遍历
func InOrder(root *Tree)  {

}


