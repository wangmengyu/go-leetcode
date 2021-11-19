package main

import (
	"fmt"
)

// 背包问题
func main() {
	// 3键物品的名称:
	productSlice := []string{"吉他", "音响", "笔记本"}
	//3件物品的价格分别是
	priceSlice := []int{1500, 3000, 2000}
	//对应的重量分别是
	weightSlice := []int{1, 4, 3}
	maxWeight := 4 // 背包总容量 4KG
	// 行数据:每一种物品. 从上到下代表可选择的商品列表
	dp := make([][]int, len(productSlice)+1)
	selectItem := make([][][]string, len(productSlice)+1)

	// 初始化表格
	for i := 0; i < len(productSlice)+1; i++ {
		dp[i] = make([]int, maxWeight+1)
		selectItem[i] = make([][]string, maxWeight+1)
		for j := 0; j < maxWeight+1; j++ {
			dp[i][j] = 0
			selectItem[i][j] = make([]string, 0)
		}
	}
	for i := 1; i <= len(productSlice); i++ {
		for j := 1; j <= maxWeight; j++ {
			// 默认值不包含当前商品
			dp[i][j] = dp[i-1][j]
			selectItem[i][j] = selectItem[i-1][j] // 沿用不包含我的选项组合
			if weightSlice[i-1] > j {
				// 当前商品重量超过了可容纳的上限, 就只能不容纳了
				continue
			}
			// 计算包含当前商品的价值
			include := dp[i-1][j-weightSlice[i-1]] + priceSlice[i-1] // 包含我:  我的价值+剩余空间的最大价值
			if include > dp[i-1][j] {                                // 当包含我比不包含我价值要高
				fmt.Printf("包含了我: %s\n", productSlice[i-1])
				dp[i][j] = include // 刷新当前单元格最高价值
				// 将包含商品设置成不包含我时的列表加上当前商品
				selectItem[i][j] = append(selectItem[i-1][j-weightSlice[i-1]], productSlice[i-1])
			}
		}
	}
	for i, row := range dp {
		fmt.Println(row)
		fmt.Println(selectItem[i])
	}
	//  输出最佳组合
	fmt.Println(selectItem[len(productSlice)][maxWeight])
}
