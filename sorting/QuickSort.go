package main

import "fmt"

func main() {
	a := []int{4, 5, 6, 1, 2, 3}
	quickSort(a, 0, len(a)-1)
	for i := range a {
		fmt.Print(a[i])
	}
}

func quickSort(a []int, start int, end int) {
	if start != end {
		pivot := partition(a, start, end)
		quickSort(a, start, pivot)
		quickSort(a, pivot+1, end)
	}
}

func partition(a []int, start int, end int) int {
	pivotValue := a[end]
	flag := start
	for j := start; j < end; j++ {
		if a[j] < pivotValue {
			if j != flag {
				a[flag], a[j] = a[j], a[flag]
				flag++
			}
		}
	}

	a[flag], a[end] = a[end], a[flag]
	return flag
}
