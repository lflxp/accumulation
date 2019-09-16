package math

import (
	"testing"
)

func Test_MergeSort(t *testing.T) {
	data := []int{8, 4, 3, 61, 7, 9, 4, 2, 21, 6, 7, 12, 5, 71}
	t.Log(data)
	MergeSort(data, 0, len(data)-1)
	t.Log(data)
}
