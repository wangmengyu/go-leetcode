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
	selectItem := make([][]map[string]int, len(productSlice)+1)

	// 初始化表格
	for i := 0; i < len(productSlice)+1; i++ {
		dp[i] = make([]int, maxWeight+1)
		selectItem[i] = make([]map[string]int, maxWeight+1)
		for j := 0; j < maxWeight+1; j++ {
			dp[i][j] = 0
			selectItem[i][j] = make(map[string]int)
		}
	}
	for i := 1; i <= len(productSlice); i++ {
		for j := 1; j <= maxWeight; j++ {
			// 当前商品的的重量是否超过的当前总重量
			if weightSlice[i-1] > j {
				//就只能不包含自己. 只能取上一行的最大值. 就是当前列的上一行
				dp[i][j] = dp[i-1][j]
				selectItem[i][j] = selectItem[i-1][j]
				continue
			}
			// 当前商品的重量没有超过当前总重量,.
			// 判断包含和不包含当前商品哪个价值大
			notInclude := dp[i-1][j]                                 // 不包含我, 上一行最大值
			include := dp[i-1][j-weightSlice[i-1]] + priceSlice[i-1] // 包含我:  我的价值+剩余空间的最大价值
			dp[i][j] = dp[i-1][j]                                    //  默认选择不包含我
			selectItem[i][j] = selectItem[i-1][j]
			if include > notInclude { // 当包含我更好,则选用包含我
				fmt.Printf("包含了我: %s\n", productSlice[i-1])
				dp[i][j] = include
				newMap := selectItem[i-1][j-weightSlice[i-1]]
				newMap[productSlice[i-1]] = 1
				selectItem[i][j] = newMap
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
