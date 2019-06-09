package main

import "fmt"

func main() {
	var a = []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	a = mergeSort(a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

}

func mergeSort(a []int) []int {
	var length = len(a)
	if length == 1 {
		return a
	}
	var mid = length / 2
	var left = a[0:mid]
	var right = a[mid:length]
	left = mergeSort(left)
	right = mergeSort(right)
	return merge2SortedArray(left, right)
}

func merge2SortedArray(a []int, b []int) []int {
	var lengthA = len(a)
	var lengthB = len(b)
	var total = lengthA + lengthB
	var count = 0
	var result = make([]int, total)
	for lengthA != 0 && lengthB != 0 {
		if a[len(a)-lengthA] < b[len(b)-lengthB] {
			result[count] = a[len(a)-lengthA]
			lengthA--
		} else {
			result[count] = b[len(b)-lengthB]
			lengthB--
		}
		count++
	}
	if lengthA != 0 {
		for lengthA != 0 {
			result[count] = a[len(a)-lengthA]
			count++
			lengthA--
		}
	}
	if lengthB != 0 {
		for lengthB != 0 {
			result[count] = b[len(b)-lengthB]
			count++
			lengthB--
		}
	}
	return result
}
