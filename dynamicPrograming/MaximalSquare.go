package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := make([][]byte, 4, 5)
	matrix[0] = []byte{1, 0, 1, 0, 0}
	matrix[1] = []byte{1, 0, 1, 1, 1}
	matrix[2] = []byte{1, 1, 1, 1, 1}
	matrix[3] = []byte{1, 0, 0, 1, 0}
	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix)+1, len(matrix[0])+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	curMax := 0
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if matrix[i-1][j-1] == 1 {
				// is adjacent points have the same area square
				left := dp[i][j-1]
				above := dp[i-1][j]
				diagonal := dp[i-1][j-1]
				if left == above && above == diagonal {
					dp[i][j] = left + 1
				} else {
					dp[i][j] = int(math.Min(math.Min(float64(left), float64(above)), float64(diagonal))) + 1
				}
				curMax = int(math.Max(float64(dp[i][j]*dp[i][j]), float64(curMax)))
			}
		}
	}
	return curMax
}
