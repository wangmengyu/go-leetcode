package main

import "fmt"

func main() {
	items := []string{
		"水",
		"书",
		"食物",
		"夹克",
		"相机",
	}
	itemNum := len(items)
	weightSlice := []int{
		3, 1, 2, 2, 1,
	}
	maxWeight := 6
	priceSlice := []int{
		10, 3, 9, 5, 6,
	}
	// 携带哪些东西价值最高
	dp := make([][]int, len(items)+1)

	dpItems := make([][][]int, len(items)+1)

	for i, _ := range dp {
		dp[i] = make([]int, maxWeight+1)
		dpItems[i] = make([][]int, maxWeight+1)

		for j, _ := range dpItems[i] {
			dpItems[i][j] = make([]int, 0)
		}
	}

	for i := 1; i <= itemNum; i++ {
		for j := 1; j <= maxWeight; j++ {
			dp[i][j] = dp[i-1][j]
			dpItems[i][j] = dpItems[i-1][j]
			//包含当前: 上一行, 除去当前物品重量的值加上物品本身的价值
			if j-weightSlice[i-1] < 0 {
				continue
			}
			include := dp[i-1][j-weightSlice[i-1]] + priceSlice[i-1]
			if include > dp[i-1][j] {
				dp[i][j] = dp[i-1][j-weightSlice[i-1]] + priceSlice[i-1]
				dpItems[i][j] = dpItems[i-1][j-weightSlice[i-1]]
				dpItems[i][j] = append(dpItems[i][j], i-1)
			}
		}
	}

	fmt.Println(dp[itemNum][maxWeight])
	fmt.Println(dpItems[itemNum][maxWeight])

}
