/**
计数排序：
	找出待排序的数组中最大和最小的元素
	统计数组中每个值为i的元素的出现次数，存入数组C的第i项
	对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
	反向填充数组：将每个元素i放在新数组的C(i)项，每放一个元素就将C(i)减去1
*/
package main

import "fmt"

func countingSort(arr []int, maxValue int) {
	bucket := make([]int, maxValue+1, maxValue+1)
	k := 0
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}
	for j := 0; j < len(bucket); j++ {
		for bucket[j] > 0 {
			arr[k] = j
			k++
			bucket[j]--
		}
	}
}

func main() {
	arr := []int{2, 9, 54, 32, 18, 29, 10, 3}
	countingSort(arr, 54)
	fmt.Println(arr)
}
