package leetcode

func removeDuplicates(nums []int) int {
	n:=len(nums)
	for i, j := 0, 1; i < n; i++ {
		if nums[i]==nums[j] {
			nums=append(nums[:i],nums[i+1:]...)
		}
	}
	return len(nums)
}
