package main

import (
	"fmt"
)

func main() {
	// 给定一个序列. 序列中抽取N个数的和要=指定的值, 有的返回true, 没有返回false
	numList := []int{3, 34, 4, 12, 5, 2}
	s := 9 // 和是9
	res := RecSubset(numList, 5, s)
	fmt.Println(res)

}

func RecSubset(numList []int, i int, s int) bool {
	if s == 0 { //
		return true
	} else if i == 0 {
		return numList[i] == s // 两个数字必须相等. 否则和不同了.
	} else if numList[i] > s {
		// 当前位置的数已经大于总数.绝对不能选, 需要跳过当前节点
		return RecSubset(numList, i-1, s)
	} else {
		// 有两种可能. 选和不选
		// 选:
		A := RecSubset(numList, i-1, s-numList[i]) // 选:的时候总值要-1
		B := RecSubset(numList, i-1, s)            // 不选: 总值不用 -1
		return A || B                              // 两种方案有一种可行就可以
	}
}

func Exists(numList []int, s int) bool {
	// 遍历每个元素. 看它选择, 还是不选择. 是能构成总和 = 9 , 可以的. 直接返回 true,  否则 , 返回false
	found := false
	opt := make([]int, len(numList))
	for i, num := range numList {
		if num == s {
			// 找到,满足了 , 直接
			found = true
			break
		} else if num > s {
			// 不能选. 已经超过了 = 上一阶段的最优值
			opt[i] = opt[i-1]
			continue
		} else if num < s {
			opt[i] = num
			// 可以选.  进行加计算
			if i-1 >= 0 {
				opt[i] += opt[i-1]
			}
			if opt[i] == 9 {
				found = true // 找到了.
				break
			}
		}
	}

	fmt.Println(opt)
	return found

}
