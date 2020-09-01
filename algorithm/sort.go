package algorithm

/**
 * 冒泡排序
 * 比较相邻元素，如果前者比后者大，则交换
 */
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

/**
 * 插入排序
 * 遍历序列，将扫描到的元素移动到适当位置
 */
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

/**
 * 选择排序
 * 在未排序序列中找到最小元素，与起始位置元素交换值
 */
func SelectSort(arr []int) []int {
	n := len(arr)
	min := 0
	for i := range arr {
		min = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

/**
 * 快速排序
 *
 */
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
	arr[obj], arr[left] = arr[left], arr[obj]
	return left
}

// 归并排序
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
