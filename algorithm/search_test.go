package algorithm

import (
	"fmt"
	"testing"
)

var arr []int

func init() {
	arr := make([]int, 1<<20)
	for i := 0; i < 1<<20; i++ {
		arr = append(arr, i)
	}
}

func TestBinarySearch(t *testing.T) {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}

func BenchmarkBinarySearch(b *testing.B) {

	for i := 0; i < b.N; i++ {
		binarySearch(arr, 1)
	}
}
