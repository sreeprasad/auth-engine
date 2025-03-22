package auth

// Wildcard match  supports
// * (match any number of characters)
// and ? (match exactly one character)
func WildcardMatch(pattern, text string) bool {
	if pattern == "" {
		return text == ""
	}
	if pattern == "*" {
		return true
	}

	dp := make([][]bool, len(pattern)+1)
	for i := range dp {
		dp[i] = make([]bool, len(text)+1)
	}

	dp[0][0] = true

	for i := 1; i <= len(pattern); i++ {
		if pattern[i-1] == '*' {
			dp[i][0] = dp[i-1][0]
			for j := 1; j <= len(text); j++ {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		} else if pattern[i-1] == '?' {
			for j := 1; j <= len(text); j++ {
				dp[i][j] = dp[i-1][j-1]
			}
		} else {
			for j := 1; j <= len(text); j++ {
				if pattern[i-1] == text[j-1] {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}

	return dp[len(pattern)][len(text)]
}
