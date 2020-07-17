package leetcode

import (
	"sort"
)

// 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	arr := make([][]int, 0)
	num := len(nums)
	for i := 0; i < num-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < num-1; j++ {
			if nums[j] > -nums[i]/2 {
				break
			}
			ij := nums[i] + nums[j]
			if ij > 0 {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < num; k++ {

				ijk := nums[k] + ij
				if ijk > 0 {
					break
				}
				if 0 == ijk {
					arr = append(arr, []int{nums[i], nums[j], nums[k]})
					break
				}
			}
		}

	}
	return arr
}


