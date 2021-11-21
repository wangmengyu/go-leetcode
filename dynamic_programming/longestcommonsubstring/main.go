package main

import "fmt"

// 最长公共子串
func main() {
	str1 := "fish"
	str2 := "hish"
	len1 := len(str1)
	len2 := len(str2)
	fmt.Println(len1)
	fmt.Println(len2)
	dp := make([][]int, len1)

	max := 0

	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
		for j := 0; j < len2; j++ {
			if str1[i] != str2[j] {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1
				if i > 0 && j > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}

	fmt.Println(max)

}
