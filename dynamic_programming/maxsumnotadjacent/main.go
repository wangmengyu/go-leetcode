package main

import (
	"fmt"
	"math"
)

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
	val := DpOpt(numList, 6)
	fmt.Println(val)
	fmt.Println(numList[6].Path)

}
func DpOpt(numList []NumItem, num int) int { // 这种算法比较快,非递归
	opt := make([]int, len(numList))
	opt[0] = numList[0].Val
	opt[1] = int(math.Max(float64(numList[0].Val), float64(numList[1].Val)))
	for i := 0; i < len(numList); i++ {
		numList[i].Path = make([]int, 0)
	}
	numList[0].Path = append(numList[0].Path, 0)
	if numList[1].Val > numList[0].Val {
		//optPath[1] = append(optPath[1], 1)
		numList[1].Path = append(numList[1].Path, 1)
	} else {
		//optPath[1] = append(optPath[1], 0)
		numList[1].Path = append(numList[1].Path, 0)
	}
	subList := numList[2:]
	for n, _ := range subList {
		n = n + 2
		fmt.Println("[n]=", n)
		A := numList[n].Val + opt[n-2] // 选 : 自己的值 + 上一个最佳值
		B := opt[n-1]                  // 不选: 上一个节点的值
		if A > B {
			opt[n] = A
			numList[n].Path = append(numList[n].Path, numList[n-2].Path...)
			numList[n].Path = append(numList[n].Path, n)
			continue
		}
		opt[n] = B

		numList[n].Path = append(numList[n].Path, numList[n-1].Path...)
	}
	fmt.Println(opt)
	return opt[num]
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
