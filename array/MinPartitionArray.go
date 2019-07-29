package main

import "math"

func main() {
	s := []float64{10, 20, 15, 5, 25}
	println(int(minPartition(s, len(s)-1, 0, 0)))
}

func minPartition(s []float64, index int, sum1 float64, sum2 float64) float64 {
	if index < 0 {
		return math.Abs(sum1 - sum2)
	}

	include := minPartition(s, index-1, sum1+s[index], sum2)
	exclude := minPartition(s, index-1, sum1, sum2+s[index])
	return math.Min(include, exclude)
}
