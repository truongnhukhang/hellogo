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
	fmt.Println("")
	heapSort(a)
	for i := 0; i < len(a); i++ {
		fmt.Print(strconv.Itoa(a[i]) + " ")
	}
}

func heapSort(a []int) {
	toMaxHeap(a)
	length := len(a)
	for i := 0; i < length; {
		swap(a, i, length-1)
		length--
		maxHeapModify(a, 1, length)
	}
}

func toMaxHeap(a []int) {
	for i := (len(a) / 2); i >= 1; i-- {
		maxHeapModify(a, i, len(a))
	}
}

func maxHeapModify(a []int, index int, length int) {
	left := left(index)
	right := right(index)
	largest := index
	if left <= length && a[left-1] > a[largest-1] {
		largest = left
	}
	if right <= length && a[right-1] > a[largest-1] {
		largest = right
	}
	if largest != index {
		swap(a, index-1, largest-1)
		maxHeapModify(a, largest, length)
	}
}

func left(index int) int {
	return 2 * index
}
func right(index int) int {
	return 2*index + 1
}

func swap(a []int, var1 int, var2 int) {
	a[var1], a[var2] = a[var2], a[var1]
}
