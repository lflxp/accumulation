// 堆排序
// 堆排序是利用堆这种数据结构而设计的一种排序算法，堆排序是一种选择排序，它的最坏，最好，平均时间复杂度均为O(nlogn)，它也是不稳定排序。
// 堆排序时间复杂度一般认为就是O(nlogn)级

// 大顶堆：每个结点的值都大于或等于其左右孩子结点的值
// 小顶堆：每个结点的值都小于或等于其左右孩子结点的值

// 根据对的特性来形成公式就是，节点为i的话
// 大顶堆: arr[i]>=arr[2i+1] && arr[i]>=arr[2i+2]
// 小顶堆：arr[i]<=arr[2i+1] && arr[i]<=arr[2i+2]

// 构建堆都是从堆的最后一个root节点开始也就是(len(arr)-1)/2
// https://blog.csdn.net/mofiu/article/details/83620743
// https://studygolang.com/articles/10896

package math

import "fmt"

func minHeap(root int, end int, c []int) {
	for {
		var child = 2*root + 1 // root为想比较的根节点、2*root+1为root的左子节点
		// 判断是否存在child节点
		if child > end {
			break
		}

		// 判断右child是否存在，如果存在则和另外一个同级节点进行比较
		if child+1 <= end && c[child] > c[child+1] {
			child += 1
		}
		if c[root] > c[child] {
			c[root], c[child] = c[child], c[root]
			root = child
		} else {
			break
		}
	}
}

// 降序排序
func HeapSort(c []int) {
	var n = len(c) - 1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, c)
	}
	fmt.Println("堆构建完成")
	for end := n; end >= 0; end-- {
		if c[0] < c[end] {
			c[0], c[end] = c[end], c[0]
			minHeap(0, end-1, c)
		}
	}
	fmt.Println(c)
}

/*
四、扩展问题
如果需要找出100w个数值中100个最大的数怎么找出呢,？
思路1: 将所有值构建大顶堆，然后排序，排序只排出100个大的值即可
思路2: 因为找出100个最大的值，不需要排序，所以可以只建一个100的数的小顶堆，然后将数组的后边的数逐一和堆顶最小值比较，如果大于堆顶的值就进行交换并重新构建堆**,如果小于顶堆的话就进行下一次比较（不重建堆）**

//在c数组中找出num个最大值
func HeapSort(c []int,num int) []int {
   m:=len(c)-1
   createHeap(c[:num],num-1)
   fmt.Println("堆构建完成")
   for i := num; i <=m; i++ {
      if c[0]<c[i]{
         c[0],c[i] = c[i],c[0]
         createHeap(c[:num],num-1)
      }
   }
   fmt.Println(c[:num])
   return c
}
func createHeap(arr []int,end int){
   for start := end / 2; start >= 0; start-- {
      minHeap(start, end, arr)
   }
}
//随机数组生成
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn((end - start))
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
————————————————
版权声明：本文为CSDN博主「誠寜」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/mofiu/article/details/83620743
*/
