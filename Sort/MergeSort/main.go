/**
归并排序：先分成2个子数组，将左右子数组排好序后合并
*/
package main

import (
	"fmt"
	"math/rand"
)

func MergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) >> 1
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)
	Merge(arr, left, mid, right)
}

func Merge(arr []int, left, mid, right int) {
	temp := make([]int, right-left+1)
	var (
		i = left
		j = mid + 1
		k = 0
	)
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = arr[i]
		k++
		i++
	}
	for j <= right {
		temp[k] = arr[j]
		k++
		j++
	}
	for p := 0; p < len(temp); p++ {
		arr[left+p] = temp[p]
	}
}

func main() {
	arr := make([]int, 10000, 10000)
	for i := 0; i < 10000; i++ {
		arr = append(arr, rand.Intn(8000))
	}
	MergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
