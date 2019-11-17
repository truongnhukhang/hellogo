package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []int{8, 1, 9, 2, 4, 5, 4, 2, 3, 10, 6, 7}
	quickSort(a, 0, len(a)-1)
	for i := range a {
		fmt.Print(" " + strconv.Itoa(a[i]))
	}
}

func quickSort(a []int, start int, end int) {
	if start < end {
		p := partition(a, start, end)
		quickSort(a, start, p-1)
		quickSort(a, p+1, end)
	}
}

func partition(a []int, start int, end int) int {
	pivot := end
	i := start
	flag := start
	for ; i < end; i++ {
		if a[i] <= a[pivot] {
			a[flag], a[i] = a[i], a[flag]
			flag++
		}
	}
	a[flag], a[pivot] = a[pivot], a[flag]
	return flag
}
