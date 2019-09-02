package main

import "fmt"

func main() {
	matrixs := []int{10, 100, 5, 50, 25}
	fmt.Println(findMinMatrixChain(matrixs))
}

func findMinMatrixChain(matrixs []int) int {
	matrixRowCol := [][]int{}
	for i := 1; i < len(matrixs); i++ {
		tempMatrix := []int{matrixs[i-1], matrixs[i]}
		matrixRowCol = append(matrixRowCol, tempMatrix)
	}
	numScalar := 0
	for i := 0; i < len(matrixRowCol); i++ {
		m1, sum1 := findMinMatrixChainByIndex(0, i, matrixRowCol)
		m2, sum2 := findMinMatrixChainByIndex(i+1, len(matrixRowCol)-1, matrixRowCol)
		tempNumScalar := sum1 + sum2 + m1[0]*m1[1]*m2[1]
		if tempNumScalar < numScalar || numScalar == 0 {
			numScalar = tempNumScalar
		}
	}
	return numScalar
}

func findMinMatrixChainByIndex(start int, end int, matrixRowCol [][]int) ([]int, int) {
	if end == start {
		return matrixRowCol[end], 0
	}
	numScalar := 0
	resultMatrix := []int{0, 0}
	for i := start; i < end; i++ {
		m1, sum1 := findMinMatrixChainByIndex(start, i, matrixRowCol)
		m2, sum2 := findMinMatrixChainByIndex(i+1, end, matrixRowCol)
		tempNumScalar := sum1 + sum2 + m1[0]*m1[1]*m2[1]
		if tempNumScalar < numScalar || numScalar == 0 {
			numScalar = tempNumScalar
			resultMatrix[0] = m1[0]
			resultMatrix[1] = m2[1]
		}
	}
	return resultMatrix, numScalar
}
