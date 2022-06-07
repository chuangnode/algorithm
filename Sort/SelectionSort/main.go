/**
  选择排序：首先在未排序序列中找到最小（大）元素，存到排序序列的起始位置，然后再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
*/
package main

import "fmt"

func SelectionSort(arr []int) []int {
	minIndex := 0
	for i := 0; i < len(arr)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	return arr
}

func main() {
	arr := []int{2, 9, 54, 32, 18, 29, 10, 3}
	fmt.Println(SelectionSort(arr))
}
