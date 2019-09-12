// 冒泡排序的基本思想就是：从无序序列头部开始，进行两两比较，根据大小交换位置，直到最后将最大（小）的数据元素交换到了无序队列的队尾，从而成为有序序列的一部分；下一次继续这个过程，直到所有数据元素都排好序。
// 算法的核心在于每次通过两两比较交换位置，选出剩余无序序列里最大（小）的数据元素放到队尾。

package math

import "fmt"

// 冒泡排序法
// https://baike.baidu.com/item/%E5%86%92%E6%B3%A1%E6%8E%92%E5%BA%8F/4602306
func BubbleSort(values []int) {
	if len(values) <= 1 {
		return
	}

	flag := true // 停止循环标签 表示左边没有数据比右边大

	for i := 0; i < len(values); i++ { // 完整数组
		fmt.Println(values)
		flag = true
		for j := 0; j < len(values)-i-1; j++ { // 可变子数组 长度为len(values)-i-1,表示排除已经确定最大数据的剩余子数组
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
}
