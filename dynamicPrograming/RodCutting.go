package main

import (
	"fmt"
	"math"
)

func main() {
	price := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	price[0] = 1
	price[1] = 5
	price[2] = 8
	price[3] = 9
	price[4] = 10
	price[5] = 17
	price[6] = 17
	price[7] = 20
	price[8] = 24
	price[9] = 30
	maxProfitMap := map[int]int{}
	fmt.Println(findMaxProfit(10, price, maxProfitMap))
}

func findMaxProfit(rodLength int, price []int, maxProfitMap map[int]int) int {
	if maxProfitMap[rodLength] != 0 {
		return maxProfitMap[rodLength]
	}
	if rodLength < 0 {
		return 0
	}
	max := price[rodLength-1]
	for i := 1; i < rodLength; i++ {
		tempMax := price[i-1] + findMaxProfit(rodLength-i, price, maxProfitMap)
		max = int(math.Max(float64(max), float64(tempMax)))
	}
	maxProfitMap[rodLength] = max
	return max
}
