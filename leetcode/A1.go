package leetcode

// 两数和
func twoSum(nums []int, target int) []int {
	m:=make(map[int]int)
	for k,v:=range nums{
		if value,ok:=m[target-v];ok {
			return []int{k,value}
		}
		m[v]=k
	}
	return nil
}
