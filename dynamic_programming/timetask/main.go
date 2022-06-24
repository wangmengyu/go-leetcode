package main

import "fmt"

type TaskItem struct {
	Start int
	End   int
	Value int
	Path  map[int]int
}

func main() {

	taskList := []TaskItem{
		{
			1, 4, 5, nil,
		},
		{
			3, 5, 1, nil,
		},
		{
			0, 6, 8, nil,
		},
		{
			4, 7, 4, nil,
		},
		{
			3, 8, 6, nil,
		},
		{
			5, 9, 3, nil,
		},
		{
			6, 10, 2, nil,
		},
		{
			8, 11, 4, nil,
		},
	}

	res := GetOpt(6, taskList)
	fmt.Println(res)
	fmt.Println(taskList[6].Path)
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
		includeVal += GetOpt(prev, taskList)
	}

	// 不要当前
	excludeVal := 0
	if n-1 >= 0 {
		// 有上一部
		excludeVal = GetOpt(n-1, taskList)
	} else {
		// 没上一部,直接返回包含当前值的公式 n + prev, 排除excluede:  n-1
		fmt.Println("[use include n]", n)
		//resMap[n] = n
		// 选择当前N + PREV的路径
		taskList[n].Path = make(map[int]int)
		if prev >= 0 {
			for k, _ := range taskList[prev].Path {
				taskList[n].Path[k] = k
			}
		}
		taskList[n].Path[n] = n

		return includeVal
	}

	if includeVal > excludeVal {
		fmt.Println("[use include n]", n)
		//resMap[n] = n

		taskList[n].Path = make(map[int]int)
		if prev >= 0 {
			for k, _ := range taskList[prev].Path {
				taskList[n].Path[k] = k
			}
		}
		taskList[n].Path[n] = n

		return includeVal
	}
	//  取exclude,  prev的path
	taskList[n].Path = taskList[n-1].Path
	fmt.Println("[use n-1]", n-1)
	return excludeVal
}
