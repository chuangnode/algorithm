package main

import (
	"fmt"
)

func BinarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	res := 0
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}

func main() {
	arr := []int{1, 4, 6, 8, 10, 20, 30}
	fmt.Println(BinarySearch(arr, 10))
}
