package math

import (
	"testing"
)

func Test_HeapSort(t *testing.T) {
	data := []int{8, 4, 3, 61, 7, 9, 4, 2, 21, 6, 7, 12, 5, 71}
	t.Log(data)
	HeapSort(data)
	t.Log(data)
}
