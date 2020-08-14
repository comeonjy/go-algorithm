package algorithm

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr:=[]int{1,2,4,7,2,4,1,8,0}
	BubbleSort(arr)
	fmt.Println(arr)
}
