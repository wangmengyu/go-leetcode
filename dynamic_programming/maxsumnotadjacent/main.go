package main

import "fmt"

type NumItem struct {
	Val  int
	Path []int
}

func main() {
	// 选取一系列数, 选择不相邻的数, 总和最大 ,
	numList := []NumItem{
		{1, nil},
		{2, nil},
		{4, nil},
		{1, nil},
		{7, nil},
		{8, nil},
		{3, nil},
	}
	res := Opt(6, numList)
	fmt.Println(res)
	fmt.Println(numList[6].Path)
}

func Opt(n int, numList []NumItem) int {

	nVal := numList[n]
	prev := Prev(n)
	if prev < 0 {
		numList[n].Path = []int{n}
		return nVal.Val // 只能选则当前值
	}
	selectNum := nVal.Val + Opt(prev, numList)
	unSelectNum := Opt(n-1, numList)
	if selectNum > unSelectNum {
		numList[n].Path = make([]int, 0)
		numList[n].Path = append(numList[n].Path, numList[prev].Path...)
		numList[n].Path = append(numList[n].Path, n)
		return selectNum
	}
	numList[n].Path = numList[n-1].Path
	return unSelectNum

}

func Prev(num int) int {
	if num-2 >= 0 {
		return num - 2
	}
	return -1
}
