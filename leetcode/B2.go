package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	head := l3
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		val := carry % 10
		carry /= 10
		l3.Next = &ListNode{Val: val}
		l3 = l3.Next
	}
	if carry>0 {
		l3.Next = &ListNode{Val: 1}
	}
	return head.Next
}
