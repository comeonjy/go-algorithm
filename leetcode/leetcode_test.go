package leetcode

import (
	"fmt"
	"testing"
)

func BenchmarkB15(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSum([]int{82597,-9243,62390,83030,-97960,-26521,-61011,83390,-38677,12333,75987,46091,83794,193})
	}
}

func TestB15(t *testing.T) {
	fmt.Println(threeSum([]int{0, 0, 0}))
}

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
