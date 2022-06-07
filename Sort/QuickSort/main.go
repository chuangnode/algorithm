/**
快速排序：通过一趟排序将待排分隔成2个独立的部分，其中一部分均比另一部分小，分别对两部分进行排序
1.从数列中跳出一个元素，作为基准（pivot）
2.重新排序序列，所以元素比基准值小的放在基准前面，所有元素比基准值大的放在后面，基准就在数列的中间位置，称为分区partition操作
3.递归地将小于基准值元素的子数列和大于基准值元素的子数列排序。
*/
package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partion(arr, left, right)
	QuickSort(arr, left, pivot-1)
	QuickSort(arr, pivot+1, right)
}

func partion(arr []int, left, right int) int {
	pivot := right
	counter := left //小于pivot的元素的个数
	for i := left; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[counter] = arr[counter], arr[i]
			counter++
		}
	}
	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return counter
}

func main() {
	arr := make([]int, 10000, 10000)
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(8000))
	}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
