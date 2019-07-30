package main

func main() {
	s := []int{4, 3, 2, 1, 6}
	println(findMissingNumber(s))
	s = []int{0, 3, 2, 1, -1}
	println(findMissingNumber(s))
}

/**
  Find missing number in an unsorted array
  Example : S = {4,3,2,1,6}
  findMissingNumber(S) = 5
*/
func findMissingNumber(s []int) int {
	min := s[0]
	max := s[0]
	sum := s[0]
	for i := 1; i < len(s); i++ {
		if min > s[i] {
			min = s[i]
		}
		if max < s[i] {
			max = s[i]
		}
		sum = sum + s[i]
	}
	total := 0
	for i := min; i <= max; i++ {
		total = total + i
	}
	missingNumber := total - sum
	if missingNumber == 0 {
		return max + 1
	}
	return missingNumber
}
