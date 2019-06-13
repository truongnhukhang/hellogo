package main

import (
	"fmt"
	. "github.com/truongnhukhang/hellogo/queue"
)

func main() {
	var a = [][]int{{1, 3, 9}, {2, 6, 10}, {4, 7, 8}, {5, 9, 11}}
	var result = mergeKSortedList(a)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

/**
merge K Sorted List
Example :
input = {1, 3, 9}, {2, 6, 10}, {4, 7, 8}, {5, 9, 11}
output = {1,2,3,4,5,6,7,8,9,9,10,11}
**/
func mergeKSortedList(a [][]int) []int {
	tempHeap := make([]int, 0)
	priorityQueue := MinPriorityQueue{tempHeap}
	tempLen := make([]int, len(a))
	result := make([]int, 0)
	max := 0
	for i := 0; i < len(a); i++ {
		tempLen[i] = len(a[i]) - 1
		priorityQueue.Put(a[i][0])
		if max < len(a[i]) {
			max = len(a[i])
		}
	}
	count := 1
	result = append(result, priorityQueue.Poll())
	for max != 0 {
		for i := 0; i < len(a); i++ {
			if tempLen[i] != 0 {
				priorityQueue.Put(a[i][count])
				result = append(result, priorityQueue.Poll())
				tempLen[i]--
			}
		}
		max--
		count++
	}
	for i := 1; i < len(a); i++ {
		result = append(result, priorityQueue.Poll())
	}
	return result
}
