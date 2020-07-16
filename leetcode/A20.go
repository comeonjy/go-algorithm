package leetcode

import (
	"go-algorithm/structure"
)

// 有效括号
func isValid(s string) bool {
	stack := structure.NewStack()
	for _, v := range s {
		switch v {
		case '[', '{', '(':
			stack.Push(v)
		case ']':
			if stack.Size()<1 || stack.Pop()!='[' {
				return false
			}
		case '}':
			if stack.Size()<1 || stack.Pop()!='{' {
				return false
			}
		case ')':
			if stack.Size()<1 || stack.Pop()!='(' {
				return false
			}
		}
	}
	return stack.Size()==0
}
