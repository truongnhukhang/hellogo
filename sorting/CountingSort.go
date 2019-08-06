package main

import "fmt"

func main() {
	s := []int{2, 5, 3, 0, 2, 3, 0, 3}
	result := countingSort(s, 5)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func countingSort(s []int, key int) []int {
	countingArray := make([]int, key+1)
	result := make([]int, len(s))
	for i := 0; i <= key; i++ {
		countingArray[i] = 0
	}
	// counting elements of original array equals with i
	for j := 0; j < len(s); j++ {
		countingArray[s[j]] = countingArray[s[j]] + 1
	}
	// counting elements of original array less than or equals with i
	for i := 1; i <= key; i++ {
		countingArray[i] = countingArray[i-1] + countingArray[i]
	}
	for j := len(s) - 1; j >= 0; j-- {
		// find the position for the element of original array in result array by counting array
		result[countingArray[s[j]]-1] = s[j]
		//Decrementing countingArray causes the next input element with a value equal to s[j],
		// if one exists, to go to the position immediately before  s[j] in the output array.
		countingArray[s[j]] = countingArray[s[j]] - 1
	}
	return result
}
