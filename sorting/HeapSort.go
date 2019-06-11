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
	var length = len(a)
	for i := 1; i <= len(a); i++ {
		maxHeapBuild(a, i)
		swapValue(a, 0, length-1)
		length--
	}
	for i := 0; i < len(a); i++ {
		fmt.Print(strconv.Itoa(a[i]) + " ")
	}
}

func maxHeapBuild(a []int, index int) {
	var left = 2 * index
	var right = 2*index + 1
	var largest = index
	if left < len(a) && a[index-1] < a[left-1] {
		largest = left
		swapValue(a, index-1, left-1)
	}
	if right < len(a) && a[index-1] < a[right-1] {
		if largest == index {
			largest = right
		}
		swapValue(a, index-1, right-1)
	}
	if largest != index {
		maxHeapBuild(a, largest)
	}
}

func swapValue(a []int, index1 int, index2 int) {
	var temp = a[index1]
	a[index1] = a[index2]
	a[index2] = temp
}
