package main

import (
	"fmt"
	. "github.com/truongnhukhang/hellogo/queue"
)

func main() {
	a := []int{3, 4, 6, 7, 1, 9}
	kLargest := findKLargestElement(a, 3)
	fmt.Println(kLargest)
}

/**
Find K'th largest element in array
input a = {3,4,6,7,1,9} , k = 2
output = 7
input a = {3,4,6,7,1,9} , k = 3
output = 6
*/
func findKLargestElement(a []int, k int) int {
	tempHeap := make([]int, 0)
	minQueue := MinPriorityQueue{tempHeap}
	count := 0
	for count != k {
		minQueue.Put(a[count])
		count++
	}
	for i := count; i < len(a); i++ {
		if a[i] > minQueue.Peek() {
			minQueue.Poll()
			minQueue.Put(a[i])
		}
	}
	return minQueue.Peek()
}
