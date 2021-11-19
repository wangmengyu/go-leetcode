package main

import "fmt"

// 二维背包问题, 求解哪些物品装入背包, 不超过 指定容量, 和指定重量. 物品的价值总和最大.

func main() {
	// 物品数量
	itemNum := 4
	// 最大容量
	maxCapacity := 5
	// 最大重量
	maxWeight := 6

	// 物品体积列表
	capacitySlice := []int{1, 2, 3, 4}

	// 物品重量
	weightSlice := []int{2, 4, 4, 5}

	// 物品价值
	priceSlice := []int{3, 4, 5, 6}

	//初始化dp数组, 比最大容量和最大重量两个维度上都+1
	dp := make([][]int, maxCapacity+1) // 一维是容量
	dpItems := make([][][]int, maxCapacity+1)
	for i, _ := range dp {
		dp[i] = make([]int, maxWeight+1) // 二维是重量
		dpItems[i] = make([][]int, maxWeight+1)
		for j, _ := range dpItems {
			dpItems[i][j] = make([]int, 0)
		}
	}

	// 遍历每个物品
	for i := 0; i < itemNum; i++ {
		//体积总大到小开始遍历, 直到遍历到当前物品的体积.
		if i > 1 {
			//break
		}
		for j := maxCapacity; j >= capacitySlice[i]; j-- {
			// 重量从大到小遍历, 直到遍历到当前物品的重量
			for k := maxWeight; k >= weightSlice[i]; k-- {
				// 选择是装入还是不装入当前物品
				notInclude := dp[j][k]
				dpItems[j][k] = dpItems[j][k]
				include := dp[j-capacitySlice[i]][k-weightSlice[i]] + priceSlice[i] // 包含我: 除了我以外的最大值 + 我的价值
				if include > notInclude {
					dp[j][k] = include // 包含我价值更高. 进行刷新
					dpItems[j][k] = append(dpItems[j-capacitySlice[i]][k-weightSlice[i]], i)
				}
			}

		}
	}
	fmt.Println(dp[maxCapacity][maxWeight])
	fmt.Println(dpItems[maxCapacity][maxWeight])

	for i, _ := range dp {
		fmt.Println(dp[i])
		fmt.Println(dpItems[i])
	}

}
