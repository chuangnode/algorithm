package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 2))
}
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartion(a, l, r)
	if q == index {
		return a[q]
	} else if q > index {
		return quickSelect(a, l, q-1, index)
	} else {
		return quickSelect(a, q+1, r, index)
	}
}

func randomPartion(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partion(a, l, r)
}

func partion(a []int, l, r int) int {
	pivot := r
	counter := l
	for i := l; i < r; i++ {
		if a[i] < a[pivot] {
			a[i], a[counter] = a[counter], a[i]
			counter++
		}
	}
	a[pivot], a[counter] = a[counter], a[pivot]
	return counter
}
