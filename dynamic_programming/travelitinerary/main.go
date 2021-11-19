package main

import "fmt"

// 旅行行程最优化问题
func main() {
	//地名
	shops := []string{
		"W",
		"G",
		"N",
		"B",
		"S",
	}

	dayNums := []int{
		1,
		1,
		2,
		4,
		1,
	}
	score := []int{
		7,
		6,
		9,
		9,
		8,
	}

	maxDay := 4
	dp := make([][]int, len(shops)+1)
	dpItems := make([][][]string, len(shops)+1)
	for i := 0; i <= len(shops); i++ {
		dp[i] = make([]int, 5)
		dpItems[i] = make([][]string, 5)
		for j := 0; j <= 4; j++ {
			dp[i][j] = 0
			dpItems[i][j] = make([]string, 0)
		}
	}

	// 遍历每个景点.
	for i := 1; i <= len(shops); i++ {
		for j := 1; j <= maxDay; j++ {
			// 默认是不选当前商店, 全部取上一行当前列的信息
			dp[i][j] = dp[i-1][j]
			dpItems[i][j] = dpItems[i-1][j]

			// 计算当前景点评分 + 除了当前景点,剩余浏览时间最高评分值
			if j-dayNums[i-1] < 0 {
				continue
			}
			include := score[i-1] + dp[i-1][j-dayNums[i-1]]

			if include > dp[i-1][j] {
				dp[i][j] = include

				dpItems[i][j] = append(dpItems[i-1][j-dayNums[i-1]], shops[i-1])

			}

		}
	}

	fmt.Println(dp[5][4])
	fmt.Println(dpItems[5][4])

}
