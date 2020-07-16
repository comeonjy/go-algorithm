package leetcode

import (
	"fmt"
	"testing"
)

func TestA21(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 5,
		},
	}
	b := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
		},
	}
	c := mergeTwoLists(a, b)
	for ; c != nil; {
		fmt.Println(c.Val)
		c = c.Next
	}
}

func TestA20(t *testing.T) {
	fmt.Println(isValid("[([])]()"))
}

func TestA14(t *testing.T) {
	fmt.Println(longestCommonPrefix([]string{
		"flower", "flow", "flight", "d",
	}))
}

func TestA13(t *testing.T) {
	fmt.Println(romanToInt("LXXIX"))
}

func TestA9(t *testing.T) {
	fmt.Println(isPalindrome(112211))
}
