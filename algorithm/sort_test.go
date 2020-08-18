package algorithm

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

var arrSort []int
var checkSort []int
func init() {
	for i := 0; i < 1000; i++ {
		arrSort = append(arrSort, rand.Intn(1000))
	}
	checkSort=make([]int,len(arrSort))
	copy(checkSort,arrSort)
	sort.Ints(checkSort)
}

func TestBubbleSort(t *testing.T) {

	BubbleSort(arrSort)
	fmt.Println(arrSort)
}

func TestInsertSort(t *testing.T) {
	InsertSort(arrSort)
	fmt.Println(arrSort)
}

func TestSelectSort(t *testing.T) {
	fmt.Println(Eq(SelectSort(arrSort),checkSort))
}

func TestQuickSort(t *testing.T) {
	fmt.Println(arrSort)
	fmt.Println(checkSort)
	sArr:=QuickSort(arrSort)
	fmt.Println(sArr)
	fmt.Println(Eq(sArr,checkSort))
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(arrSort)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(arrSort)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(arrSort)
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort(arrSort)
	}
}


func Eq(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}
