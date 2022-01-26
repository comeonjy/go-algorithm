package algorithm

// BinarySearch 二分查找
func BinarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == arr[mid] {
			return mid
		}
		if target > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
