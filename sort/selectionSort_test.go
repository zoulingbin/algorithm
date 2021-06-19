package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 5, 6, 7, 8, 3, 2, 1}
	selectionSort(arr)
	fmt.Println(arr)
}