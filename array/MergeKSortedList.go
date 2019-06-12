package main

import (
	"fmt"
	. "github.com/truongnhukhang/hellogo/heap"
)

func main() {
	var a = [][]int{{1, 3, 9}, {2, 6, 10}, {4, 7, 8}, {5, 9, 11}}
	var result = mergeKSortedList(a)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func mergeKSortedList(a [][]int) []int {
	if len(a) == 1 {
		return a[0]
	}
	mid := len(a) / 2
	aLeft := mergeKSortedList(a[0:mid])
	aRight := mergeKSortedList(a[mid:len(a)])
	result := make([]int, len(aLeft)+len(aRight))
	length := len(aLeft)
	for i := 0; i < len(aLeft); i++ {
		result[i] = aLeft[i]
	}
	for i := 0; i < len(aRight); i++ {
		InsertToMinHeap(result, aRight[i], length+1)
		length++
	}
	return result
}
