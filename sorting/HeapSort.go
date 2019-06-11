package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	//maxHeapBuild(a,1);
	//for i:=0;i<len(a);i++ {
	//	fmt.Print(strconv.Itoa(a[i]) + " ");
	//}
	sortInplace(a)
	for i := 0; i < len(a); i++ {
		fmt.Print(strconv.Itoa(a[i]) + " ")
	}
}

func sortInplace(a []int) {
	toMaxHeap(a)
	var length = len(a)
	for i := length; i > 1; i-- {
		swapValue(a, 0, i-1)
		length--
		maxHeapBuild(a, 1, length)
	}
}

func sort(a []int) []int {
	var length = len(a)
	var result []int
	for length != 0 {
		maxHeapBuild(a, 1, len(a))
		result = append(result, a[0])
		a = a[1:len(a)]
		length--
	}
	return result
}
func toMaxHeap(a []int) {
	for i := len(a) / 2; i > 0; i-- {
		maxHeapBuild(a, i, len(a))
	}
}
func maxHeapBuild(a []int, index int, length int) {
	var left = 2 * index
	var right = 2*index + 1
	var largest = index
	if left <= length && a[index-1] < a[left-1] {
		largest = left
	}
	if right <= length && a[largest-1] < a[right-1] {
		largest = right
	}
	if largest != index {
		swapValue(a, index-1, largest-1)
		maxHeapBuild(a, largest, length)
	}
}

func swapValue(a []int, index1 int, index2 int) {
	var temp = a[index1]
	a[index1] = a[index2]
	a[index2] = temp
}
