package leetcode

// 移动零
func moveZeroes1(nums []int) {
	index := 0

	for _, v := range nums {
		if v != 0 {
			nums[index] = v
			index++
		}
	}
	for ; index < len(nums); index++ {
		nums[index] = 0
	}

}

func moveZeroes(nums []int) {
	n := len(nums)
	for i, j := 0, 0; i < n; i++ {
		if nums[i] != 0   {
			if i>j {
				nums[j], nums[i] = nums[i], nums[j]
			}
			j++
		}
	}
}

func moveZeroes2(nums []int) {
	n := len(nums)
	for i, j := 0, 0; j < n; i++ {
		if i < n {
			if nums[i] != 0 {
				nums[j] = nums[i]
				j++
			}
		} else {
			nums[j] = 0
			j++
		}
	}
}
