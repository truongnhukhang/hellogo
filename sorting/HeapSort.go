package main

import (
	"fmt"
	. "github.com/truongnhukhang/hellogo/heap"
	"strconv"
)

func main() {
	var a = []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1, 18}
	//sort(a);
	ToMinHeap(a)
	for i := 0; i < len(a); i++ {
		fmt.Print(strconv.Itoa(a[i]) + " ")
	}
}

func sort(a []int) {
	ToMaxHeap(a)
	var length = len(a)
	for i := length; i > 1; i-- {
		swapValue(a, 0, i-1)
		length--
		MaxHeapBuild(a, 1, length)
	}
}

func swapValue(a []int, index1 int, index2 int) {
	var temp = a[index1]
	a[index1] = a[index2]
	a[index2] = temp
}
