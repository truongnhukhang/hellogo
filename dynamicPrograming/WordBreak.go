package main

func main() {
	wordDict := []string{"kh", "ng", "ang"}
	println(wordBreak("khang", wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	sMap := make(map[string]bool)
	for i := 0; i < len(wordDict); i++ {
		sMap[wordDict[i]] = true
	}
	dp := make([]bool, len(s))
	dp[0] = sMap[s[0:1]]
	for i := 1; i < len(s); i++ {
		for j := i; j >= 0; j-- {
			sub := s[j : i+1]
			if sMap[sub] && j != 0 {
				dp[i] = dp[j-1]
				if dp[i] {
					break
				}
			}
			if sMap[sub] && j == 0 {
				dp[i] = true
			}
		}
	}
	return dp[len(s)-1]
}
