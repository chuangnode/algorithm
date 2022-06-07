/**
希尔排序：简单插入排序的改进版，优先比较距离远的元素，又叫缩小增量排序
将整个待排序的记录序列分割成为若干个子序列进行插入排序。
*/
package main

import (
	"fmt"
	"math"
)

func ShellSort(arr []int) []int {
	var j, cur int
	for gap := int(math.Floor(float64(len(arr) / 2))); gap > 0; gap = int(math.Floor(float64(gap / 2))) {
		for i := gap; i < len(arr); i++ {
			j = i
			cur = arr[i]
			for j-gap >= 0 && arr[j-gap] > cur {
				arr[j] = arr[j-gap]
				j = j - gap
			}
			arr[j] = cur
		}
	}
	return arr
}

func main() {
	arr := []int{2, 9, 54, 32, 18, 29, 10, 3}
	fmt.Println(ShellSort(arr))
}
