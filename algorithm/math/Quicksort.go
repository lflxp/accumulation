// 快速排序（Quicksort）是对冒泡排序的一种改进。 [1]
// 快速排序由C. A. R. Hoare在1960年提出。它的基本思想是：
// 通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，
// 然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。
package math

import "fmt"

/*
排序流程
快速排序算法通过多次比较和交换来实现排序，其排序流程如下： [2]
(1)首先设定一个分界值，通过该分界值将数组分成左右两部分。 [2]
(2)将大于或等于分界值的数据集中到数组右边，小于分界值的数据集中到数组的左边。此时，左边部分中各元素都小于或等于分界值，而右边部分中各元素都大于或等于分界值。 [2]
(3)然后，左边和右边的数据可以独立排序。对于左侧的数组数据，又可以取一个分界值，将该部分数据分成左右两部分，同样在左边放置较小值，右边放置较大值。右侧的数组数据也可以做类似处理。 [2]
(4)重复上述过程，可以看出，这是一个递归定义。通过递归将左侧部分排好序后，再递归排好右侧部分的顺序。当左、右两个部分各数据排序完成后，整个数组的排序也就完成了。 [2]

[879 637 408 984 896 736 521 999 905 151]
i = 1 values[i] = 637  mid = 879  head = 0 tail = 9
[637 879 408 984 896 736 521 999 905 151]
i = 2 values[i] = 408  mid = 879  head = 1 tail = 9
[637 408 879 984 896 736 521 999 905 151]
i = 3 values[i] = 984  mid = 879  head = 2 tail = 9
[637 408 879 151 896 736 521 999 905 984]
i = 3 values[i] = 151  mid = 879  head = 2 tail = 8
[637 408 151 879 896 736 521 999 905 984]
i = 4 values[i] = 896  mid = 879  head = 3 tail = 8
[637 408 151 879 905 736 521 999 896 984]
i = 4 values[i] = 905  mid = 879  head = 3 tail = 7
[637 408 151 879 999 736 521 905 896 984]
i = 4 values[i] = 999  mid = 879  head = 3 tail = 6
[637 408 151 879 521 736 999 905 896 984]
i = 4 values[i] = 521  mid = 879  head = 3 tail = 5
[637 408 151 521 879 736 999 905 896 984]
i = 5 values[i] = 736  mid = 879  head = 4 tail = 5

静态值： mid
变量: i 和 tail、head以及i指向的值，该值有两个状态：一个是i index位置的值，一个是比对mid后的替换值
*/

// 排序步骤（重点3、4步骤 先从后往前再从前往后）
// 原理
// 设要排序的数组是A[0]……A[N-1]，首先任意选取一个数据作为关键数据，然后将所有比它小的数都放到它左边，所有比它大的数都放到它右边，这个过程称为一趟快速排序。值得注意的
// 一趟快速排序的算法是： [1]
// 1）设置两个变量i、j，排序开始的时候：i=0，j=N-1； [1]
// 2）以第一个数组元素作为关键数据，赋值给key，即key=A[0]； [1]
// 3）从j开始向前搜索，即由后开始向前搜索(j--)，找到第一个小于key的值A[j]，将A[j]和A[i]的值交换； [1]
// 4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]和A[j]的值交换； [1]
// 5）重复第3、4步，直到i=j； (3,4步中，没找到符合条件的值，即3中A[j]不小于key,4中A[i]不大于key的时候改变j、i的值，使得j=j-1，i=i+1，直至找到为止。找到符合条件的值，进行交换的时候i， j指针位置不变。另外，i==j这一过程一定正好是i+或j-完成的时候，此时令循环结束）。 [1]
func QuickSort(values []int) {
	if len(values) <= 1 {
		return
	}

	// 获取数组第一个数据mid作为关键数据
	// 取index=1，即数组第二个数据作为开始比对数据
	mid, i := values[0], 1
	head, tail := 0, len(values)-1 //定义比对数据的边界

	for head < tail { // 遍历数组左右元素 O(log2n)
		fmt.Println(values)
		// values[i] i第一次循环的时候一直不变，都是将tail移到i的位置，然后tail前移
		fmt.Printf("i = %d values[i] = %d  mid = %d  head = %d tail = %d\n", i, values[i], mid, head, tail)
		if values[i] > mid { // 把大于mid的数据统统放到 mid右边 values[i]
			values[i], values[tail] = values[tail], values[i]
			tail-- // 3）从j开始向前搜索，即由后开始向前搜索(j--)，找到第一个小于key的值A[j]，将A[j]和A[i]的值交换； [1]
		} else { // 把小于mid的数据放到 mid右边 mid会位移到中位区（i位置没变，但是数据一直都由tail数据提供）
			values[i], values[head] = values[head], values[i]
			head++ // 4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]和A[j]的值交换
			i++    // 左边排满了 移动一个位置存后续小于mid的数
		}
	}
	values[head] = mid         // 将中间head值替换为mid
	QuickSort(values[:head])   // 递归排序右边数据
	QuickSort(values[head+1:]) // 递归排序左边数据
}
