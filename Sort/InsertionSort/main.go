/**
插入排序：
	1.从第一个元素开始，该元素可以被认为已经排序
	2.取出下一个元素，在已经排序的元素序列中从后向前扫描
	3.如果该元素（已排序）大于新元素，将该元素移到下一位置
	4.重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
	5.将新元素插入到该位置后
	6.重复步骤2-5
*/
package main

import "fmt"

func InsertionSort(arr []int) []int {
	var prev, cur int
	for i := 1; i < len(arr); i++ {
		cur = arr[i]
		prev = i - 1
		for prev >= 0 && arr[prev] > cur {
			arr[prev+1] = arr[prev]
			prev--
		}
		arr[prev+1] = cur
	}
	return arr
}

func main() {
	arr := []int{2, 9, 54, 32, 18, 29, 10, 3}
	fmt.Println(InsertionSort(arr))
}
