package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findLongestCommonSubsequence("ABCBDABD", "BDCABACD"))
}

func findLongestCommonSubsequence(s1 string, s2 string) string {
	s1L := len(s1)
	s2L := len(s2)
	s1Array := strings.Split(s1, "")
	s2Array := strings.Split(s2, "")
	path := make([][]int, s1L+1)
	direction := make([][]string, s1L+1)
	for i := 0; i < len(path); i++ {
		path[i] = make([]int, s2L+1)
		direction[i] = make([]string, s2L+1)
	}
	for i := 0; i < len(path); i++ {
		path[i][0] = 0
	}
	for i := 0; i < len(path[0]); i++ {
		path[0][i] = 0
	}
	for i := 1; i <= s1L; i++ {
		for j := 1; j <= s2L; j++ {
			if s1Array[i-1] == s2Array[j-1] {
				path[i][j] = path[i-1][j-1] + 1
				direction[i][j] = "d"
			} else {
				if path[i-1][j] >= path[i][j-1] {
					path[i][j] = path[i-1][j]
					direction[i][j] = "u"
				} else {
					path[i][j] = path[i][j-1]
					direction[i][j] = "l"
				}
			}
		}
	}
	for i := 0; i <= s1L; i++ {
		for j := 0; j <= s2L; j++ {
			fmt.Print(path[i][j])
		}
		fmt.Println("")
	}
	result := ""
	buildLongestCommonSubsequenceByMatrix(direction, s1L, s2L, s2Array, &result)
	return result
}

func buildLongestCommonSubsequenceByMatrix(direction [][]string, x int, y int, source []string, result *string) {
	if x >= 0 && y >= 0 {
		if direction[x][y] == "d" {
			buildLongestCommonSubsequenceByMatrix(direction, x-1, y-1, source, result)
			*result = *result + source[y-1]
		} else if direction[x][y] == "u" {
			buildLongestCommonSubsequenceByMatrix(direction, x-1, y, source, result)
		} else {
			buildLongestCommonSubsequenceByMatrix(direction, x, y-1, source, result)
		}
	}
}
