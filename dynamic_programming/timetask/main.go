package main

import "fmt"

type TaskItem struct {
	Start int
	End   int
	Value int
}

func main() {
	taskList := []TaskItem{
		{
			1, 4, 5,
		},
		{
			3, 5, 1,
		},
		{
			0, 6, 8,
		},
		{
			4, 7, 4,
		},
		{
			3, 8, 6,
		},
		{
			5, 9, 3,
		},
		{
			6, 10, 2,
		},
		{
			8, 11, 4,
		},
	}

	res := GetOpt(7, taskList)
	fmt.Println(res)
}

func Prev(n int, taskList []TaskItem) int {
	//获得最近的上一步可操作节点的位置
	if n == 0 {
		return -1 // 代表没有
	}
	currPos := taskList[n].Start
	for start := n - 1; start >= 0; start-- {
		// 逐个找,有结束时间在开始时间之前的就返回.否则返回-1

		preEnd := taskList[start].End
		if preEnd <= currPos {
			return start
		}
	}
	return -1
}

func GetOpt(n int, taskList []TaskItem) int {
	//  要当前
	prev := Prev(n, taskList)
	includeVal := taskList[n].Value
	if prev > -1 {
		fmt.Println(prev)
		includeVal += GetOpt(prev, taskList)
	}

	// 不要当前
	excludeVal := 0
	if n-1 >= 0 {
		excludeVal = GetOpt(n-1, taskList)
	} else {
		return includeVal
	}

	fmt.Println("include:", includeVal)
	fmt.Println("exclude:", excludeVal)
	if includeVal > excludeVal {
		return includeVal
	}
	return excludeVal
}
