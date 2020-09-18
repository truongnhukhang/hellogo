package main

import "strconv"

func main() {
	println(numDecodings("12012"))
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		temp := s[0:1]
		intVal, _ := strconv.ParseInt(temp, 10, 32)
		if intVal == 0 {
			return 0
		} else {
			return 1
		}
	}
	if len(s) > 1 {
		temp := s[0:1]
		intVal, _ := strconv.ParseInt(temp, 10, 32)
		if intVal == 0 {
			return 0
		}
	}
	dp := make([]int, len(s))
	dp[0] = 1
	dp[1] = 0
	temp, _ := strconv.ParseInt(s[1:2], 10, 32)
	temp2num, _ := strconv.ParseInt(s[0:2], 10, 32)
	if (temp2num >= 1 && temp2num <= 26) || (temp != 0) {
		dp[1] = 1
	}
	if temp2num >= 1 && temp2num <= 26 && temp != 0 {
		dp[1] = 2
	}
	for i := 2; i < len(s); i++ {
		prvVal, _ := strconv.ParseInt(s[i-1:i], 10, 32)
		temp, _ := strconv.ParseInt(s[i:i+1], 10, 32)
		temp2num, _ := strconv.ParseInt(s[i-1:i+1], 10, 32)
		if temp != 0 {
			dp[i] = dp[i-1]
		}
		if prvVal != 0 && temp2num >= 1 && temp2num <= 26 {
			dp[i] = dp[i-2]
		}
		if prvVal != 0 && temp2num >= 1 && temp2num <= 26 && temp != 0 {
			dp[i] = dp[i-2] + dp[i-1]
		}
	}
	return dp[len(s)-1]
}
