package algorithm_test

import (
	"testing"

	"go-algorithm/algorithm"
	"go-algorithm/pkg/xtest"
)

var arr []int

func init() {
	arr = make([]int, 0)
	for i := 0; i < 1<<20; i++ {
		arr = append(arr, i)
	}
}

func TestBinarySearch(t *testing.T) {
	xtest.Assert(t, algorithm.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5), 4)
	xtest.Assert(t, algorithm.BinarySearch([]int{0, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0), 0)
	xtest.Assert(t, algorithm.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10), 9)
}

func BenchmarkBinarySearch(b *testing.B) {
	b.Run("good", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			algorithm.BinarySearch(arr, 1)
		}
	})
}
