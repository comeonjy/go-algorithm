package leetcode


// 颜色分类
func sortColors(nums []int)  {
	for i,zero,two:=0,0,len(nums)-1;i<two; {
		switch nums[i] {
		case 0:
			nums[zero],nums[i]=nums[i],nums[zero]
			zero++
			i++
		case 1:
			i++
		case 2:
			nums[two],nums[i]=nums[i],nums[two]
			two--
		}
	}
}