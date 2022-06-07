/**
桶排序：
	设置一个定量的数组当作空桶
	遍历输入数据，并且把数据一个一个放到对应的桶里去
	对每个不是空的桶进行排序
	从不是空的桶里把排好序的数据拼接起来
*/
package main

import (
	"fmt"
	"math"
)

func InsertionSort(arr []int) {
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
}
func BucketSort(arr []int, bucketSize int) {
	var minValue, maxValue = arr[0], arr[0]
	buckCount := int(math.Floor(float64((maxValue-minValue)/bucketSize))) + 1
	buckets := make([][]int, buckCount)
	index := 0
	for i := 0; i < len(arr); i++ {
		index = int(math.Floor(float64((arr[i] - minValue) / bucketSize)))
		buckets[index] = append(buckets[index], arr[i])
	}
	for i := 0; i < len(buckets); i++ {
		InsertionSort(buckets[i])
		for j := 0; j < len(buckets[i]); j++ {
			arr[j] = buckets[i][j]
		}
	}
}

func main() {
	arr := []int{2, 9, 54, 32, 18, 29, 10, 3}
	BucketSort(arr, 54)
	fmt.Println(arr)
}
