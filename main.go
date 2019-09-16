package main

import (
	"fmt"
	"math/rand"

	"github.com/lflxp/accumulation/algorithm/math"
)

func randInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func main() {
	// test := reverselinkedlist.NewList()
	// reverse := reverselinkedlist.ReverseList(test)
	// fmt.Println(reverse)

	data := []int{}
	for i := 0; i < 10; i++ {
		data = append(data, randInt(1, 1000))
	}

	math.QuickSort(data)
	// math.BubbleSort(data)
	fmt.Println(data)

	// data := []int{8, 4, 3, 61, 7, 9, 4, 2, 21, 6, 7, 12, 5, 71}
	// fmt.Println(data)
	// math.MergeSort(data, 0, len(data))
	// fmt.Println(data)
}
