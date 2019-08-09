package main

import "fmt"

func main() {
	s := []int{329, 457, 657, 839, 436, 720, 355}
	result := radixSort(s, 3)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func radixSort(s []int, numDigit int) []int {
	for digit := 1; digit <= numDigit; digit++ {
		tempResult := make([]int, len(s))
		counting := make([]int, 10)
		for i := 0; i < 10; i++ {
			counting[i] = 0
		}
		for j := 0; j < len(s); j++ {
			counting[getDigitAtIndex(s[j], digit)] = counting[getDigitAtIndex(s[j], digit)] + 1
		}
		for i := 1; i < len(counting); i++ {
			counting[i] = counting[i-1] + counting[i]
		}
		for i := len(s) - 1; i >= 0; i-- {
			tempResult[counting[getDigitAtIndex(s[i], digit)]-1] = s[i]
			counting[getDigitAtIndex(s[i], digit)]--
		}
		s = tempResult
	}
	return s

}

func getDigitAtIndex(num int, index int) int {
	count := 0
	result := -1
	for count != index && num != 0 {
		count++
		result = num % 10
		num = num / 10
	}
	return result
}
