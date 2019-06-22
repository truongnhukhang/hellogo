package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(search(a, 5))
}

func search(a []int, value int) int {
	return binarySearch(a, 0, len(a)-1, value)
}

func binarySearch(a []int, left int, right int, searchValue int) int {
	if left == right {
		return left
	}
	mid := math.Floor(float64((right + left) / 2))
	if searchValue < a[int(mid)] {
		return binarySearch(a, left, int(mid)-1, searchValue)
	} else if searchValue > a[int(mid)] {
		return binarySearch(a, int(mid)+1, right, searchValue)
	}
	return -1
}
