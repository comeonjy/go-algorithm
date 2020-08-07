package leetcode

import (
	"go-algorithm/structure"
)

// 下一个更大元素
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	stack := structure.NewStack()
	for _, v := range nums2 {
		if stack.Head != nil && stack.Head.Data.(int) < v {
			for stack.Head != nil {
				if stack.Head.Data.(int) >= v {
					break
				} else {
					m[stack.Pop().(int)] = v
				}
			}
		}
		stack.Push(v)
	}
	for k, v := range nums1 {
		if value, ok := m[v]; ok {
			nums1[k] = value
		} else {
			nums1[k] = -1
		}
	}
	return nums1
}
