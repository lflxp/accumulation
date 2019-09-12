package math

import (
	"math/rand"
	"testing"
)

func randInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func Test_QuickSort(t *testing.T) {
	data := []int{}
	for i := 0; i < 100; i++ {
		data = append(data, randInt(1, 100))
	}

	QuickSort(data)
	t.Log(data)
}
