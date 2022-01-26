package algorithm

// 排序算法

// BubbleSort 冒泡排序
// 原理：比较相邻元素，如果前者比后者大，则交换
// 时间复杂度：O(n2)
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

// InsertSort 插入排序
// 原理：遍历序列，将扫描到的元素移动到适当位置
// 时间复杂度：O(n2)
func InsertSort(arr []int) []int {
	for i, v := range arr {
		j := i - 1
		for ; j >= 0 && v < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = v
	}
	return arr
}

// SelectSort 选择排序
// 原理：在未排序序列中找到最小元素，与起始位置元素交换值
// 时间复杂度：O(n2)
func SelectSort(arr []int) []int {
	n := len(arr)
	minIndex := 0
	for i := range arr {
		minIndex = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// QuickSort 快速排序
// 原理：找一个基点，使其左右两边有序，然后分别对其左右两部分进行排序
// 时间复杂度：O(nlogn)
func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left, right int) []int {
	if left < right {
		index := partition(arr, left, right)
		quickSort(arr, left, index-1)
		quickSort(arr, index+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	obj := left
	for left < right {
		if arr[right] > arr[obj] {
			right--
		} else {
			if arr[left] <= arr[obj] {
				left++
			} else {
				arr[left], arr[right] = arr[right], arr[left]
				right--
			}
		}
	}
	arr[obj], arr[right] = arr[right], arr[obj]
	return left
}

// MergeSort 归并排序
// 原理：将数组拆分为很多小数组分别排序，两两合并（适用于大量数据进行分布式排序）
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	return merge(MergeSort(arr[0:middle]), MergeSort(arr[middle:]))
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, 0)
	for len(arr1) != 0 && len(arr2) != 0 {
		if arr1[0] < arr2[0] {
			result = append(result, arr1[0])
			arr1 = arr1[1:]
		} else {
			result = append(result, arr2[0])
			arr2 = arr2[1:]
		}
	}
	if len(arr1) > 0 {
		result = append(result, arr1...)
	}
	if len(arr2) > 0 {
		result = append(result, arr2...)
	}
	return result
}
