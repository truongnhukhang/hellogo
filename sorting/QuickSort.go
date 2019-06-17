package main

import "fmt"

func main() {
	a := []int{8, 1, 9, 2, 4, 5, 4, 2, 3, 10, 6, 7}
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
			a[flag], a[j] = a[j], a[flag]
			flag++
		}
	}
	if flag != end {
		a[flag], a[end] = a[end], a[flag]
		return flag
	} else {
		return flag - 1
	}
}
