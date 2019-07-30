package main

import "math"

func main() {
	s := []float64{10, 20, 15, 5, 25}
	println(int(minPartition(s, len(s)-1, 0, 0)))
}

/**
	Find min sub-set array
    give array s = {10,20,15,5,25}
	return minimum distance of 2 sum subset
    f(s) = 5 . Because Sum({25,15})-Sum({10,15,5}) = 5 .
*/
func minPartition(s []float64, index int, sum1 float64, sum2 float64) float64 {
	if index < 0 {
		return math.Abs(sum1 - sum2)
	}
	include := minPartition(s, index-1, sum1+s[index], sum2)
	exclude := minPartition(s, index-1, sum1, sum2+s[index])
	return math.Min(include, exclude)
}
