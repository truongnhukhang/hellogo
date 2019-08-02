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
	listNumberArray := len(a)
	result := []int{}
	maxlength := len(a[0])
	// init minPriorityQueue and find Max length Array;
	minQueue := MinPriorityQueue{[]int{}}
	for i := 0; i < listNumberArray; i++ {
		minQueue.Put(a[i][0])
		if len(a[i]) > maxlength {
			maxlength = len(a[i])
		}
	}

	// start merge . We loop for each element in each list . Put it into PriorityQueue
	// Poll() the minimum element out of the Queue and append to the result.
	count := 1
	for count != maxlength {
		for i := 0; i < listNumberArray; i++ {
			result = append(result, minQueue.Poll())
			if count < len(a[i]) {
				minQueue.Put(a[i][count])
			}
		}
		count++
	}

	for !minQueue.IsEmpty() {
		result = append(result, minQueue.Poll())
	}
	return result
}
