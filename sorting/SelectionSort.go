package main

import "fmt"

func main() {
	var a = []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1, 18}
	selectionSort(a)
	for i := range a {
		fmt.Println(a[i])
	}
}

func selectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}
