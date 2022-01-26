package algorithm_test

import (
	"math/rand"
	"sort"
	"testing"

	"go-algorithm/algorithm"
	"go-algorithm/pkg/xtest"
)

var arrSort []int
var checkSort []int

func init() {
	for i := 0; i < 10000; i++ {
		arrSort = append(arrSort, rand.Intn(10000))
	}
	checkSort = make([]int, len(arrSort))
	copy(checkSort, arrSort)
	sort.Ints(checkSort)
}

func TestSort(t *testing.T) {
	a := make([]int, len(arrSort))
	t.Run("QuickSort", func(t *testing.T) {
		copy(a, arrSort)
		xtest.Assert(t, algorithm.QuickSort(a), checkSort)
	})
	t.Run("MergeSort", func(t *testing.T) {
		copy(a, arrSort)
		xtest.Assert(t, algorithm.MergeSort(a), checkSort)
	})
	t.Run("SelectSort", func(t *testing.T) {
		copy(a, arrSort)
		xtest.Assert(t, algorithm.SelectSort(a), checkSort)
	})
	t.Run("InsertSort", func(t *testing.T) {
		copy(a, arrSort)
		xtest.Assert(t, algorithm.InsertSort(a), checkSort)
	})
	t.Run("BubbleSort", func(t *testing.T) {
		copy(a, arrSort)
		xtest.Assert(t, algorithm.BubbleSort(a), checkSort)
	})

}

func BenchmarkSort(b *testing.B) {
	a := make([]int, len(arrSort))

	b.Run("QuickSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(a, arrSort)
			algorithm.QuickSort(a)
		}
	})

	b.Run("sort.Ints", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(a, arrSort)
			sort.Ints(a)
		}
	})

	b.Run("MergeSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(a, arrSort)
			algorithm.MergeSort(a)
		}
	})

	b.Run("InsertSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(a, arrSort)
			algorithm.InsertSort(a)
		}
	})

	b.Run("SelectSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(a, arrSort)
			algorithm.SelectSort(a)
		}
	})

	b.Run("BubbleSort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			copy(a, arrSort)
			algorithm.BubbleSort(a)
		}
	})

}
