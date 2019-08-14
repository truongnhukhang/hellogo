package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []float64{0.78, 0.17, 0.89, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}
	result := bucketSort(s)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func bucketSort(s []float64) []float64 {
	result := make([]float64, 0, len(s))
	length := len(s)
	bucketArray := make([][]float64, len(s))
	for i := 0; i < len(bucketArray); i++ {
		bucketArray[i] = make([]float64, 0, len(s))
	}
	for i := 0; i < len(s); i++ {
		bucketIndex := int(s[i]*float64(length)) - 1
		bucketArray[bucketIndex] = append(bucketArray[bucketIndex], s[i])
	}
	for i := 0; i < len(bucketArray); i++ {
		sort.Float64s(bucketArray[i])
	}
	for i := 0; i < len(bucketArray); i++ {
		if len(bucketArray[i]) > 0 {
			result = append(result, bucketArray[i]...)
		}
	}
	return result
}
