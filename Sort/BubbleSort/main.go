/**
冒泡排序是一种简单的排序算法，一次比较2个元素，如果顺序错误就交换，一直到没有再需要交换。
*/
package main

import "fmt"

func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{2, 9, 54, 32, 18, 29, 10, 3}
	fmt.Println(BubbleSort(arr))
}
