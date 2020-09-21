package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(13))
}

func numSquares(n int) int {
	cache := make(map[int]*int)
	return findMind(n, cache)
}

func findMind(n int, cache map[int]*int) int {
	if n == 0 {
		return 0
	}
	if cache[n] != nil {
		return *cache[n]
	}
	sqrtN := int(math.Sqrt(float64(n)))
	min := math.MaxInt32
	for i := 1; i <= sqrtN; i++ {
		min = int(math.Min(float64(min), float64(1+findMind(n-i*i, cache))))
	}
	cache[n] = &min
	return min
}
