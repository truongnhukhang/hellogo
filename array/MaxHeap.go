package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1, 18}
	toMaxHeap(a)
	for i := 0; i < len(a); i++ {
		fmt.Print(strconv.Itoa(a[i]) + " ")
	}
}

func toMaxHeap(a []int) {
	for i := (len(a) / 2); i >= 1; i-- {
		maxHeapModify(a, i)
	}
}

func maxHeapModify(a []int, index int) {
	left := left(index)
	right := right(index)
	largest := index
	if a[left-1] > a[largest-1] {
		swap(a, largest, left)
		largest = left
	}
	if a[right-1] > a[largest-1] && largest != left {
		swap(a, largest, right)
		largest = right
	}
	if largest != index {
		maxHeapModify(a, largest)
	}
}

func left(index int) int {
	return 2 * index
}
func right(index int) int {
	return 2*index + 1
}

func swap(a []int, var1 int, var2 int) {
	a[var1-1], a[var2-1] = a[var2-1], a[var1-1]
}
