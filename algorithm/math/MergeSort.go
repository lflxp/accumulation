// 归并排序(核心思想：分治)
// 归并排序的操作步骤如下：

// 首先将数组一份为二，分别为左数组和右数组
// 若左数组的长度大于1，那么对左数组实施归并排序
// 若右数组的长度大于1， 那么对右数组实施归并排序
// 将左右数组进行合并
// https://blog.csdn.net/k_koris/article/details/80508543
// https://blog.csdn.net/qq_36520153/article/details/82894216 有图

package math

import "fmt"

func MergeSort(arr []int, low, heigh int) {
	if heigh-low <= 1 {
		return
	}

	mid := (low + heigh) / 2
	// fmt.Println(len(arr), low, mid, heigh)
	MergeSort(arr, low, mid)
	MergeSort(arr, mid, heigh)

	arrLeft := make([]int, mid-low)
	arrRight := make([]int, heigh-mid)
	copy(arrLeft, arr[low:mid])
	copy(arrRight, arr[mid:heigh])

	Lindex := 0
	Rindex := 0

	for k := low; k < heigh; k++ {
		fmt.Println("arrLeft", arrLeft, arrRight, low, mid, heigh, Lindex, Rindex, len(arr), k)
		if Lindex >= mid-low { // 排序
			fmt.Println("Lindex", Lindex, mid-low, arrRight[Rindex], Rindex)
			arr[k] = arrRight[Rindex]
			fmt.Println("排序 Rindex", arr)
			Rindex++
		} else if Rindex >= heigh-mid { // 排序
			fmt.Println("Rindex", Rindex, heigh-mid, arrLeft[Lindex], Lindex)
			arr[k] = arrLeft[Lindex]
			fmt.Println("排序 Lindex", arr)
			Lindex++
		} else if arrLeft[Lindex] < arrRight[Rindex] { // 归并 获取小的
			fmt.Println("归并前 L", k, arr)
			arr[k] = arrLeft[Lindex]
			fmt.Println("归并后 L", k, arr)
			Lindex++
		} else {
			fmt.Println("归并前 R", k, arr)
			arr[k] = arrRight[Rindex] // 归并 获取小的
			fmt.Println("归并后 R", k, arr)
			Rindex++
		}
	}
}
