package main

import "fmt"

/**
按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n皇后问题 研究的是如何将 n个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的n皇后问题 的解决方案。

每一种解法包含一个不同的n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

*/
func main() {
	val := totalNQueens(8)
	fmt.Println(val)
}
func isNotUnderAttack(row, col, n int, rows, hills, dales []int) bool {
	res := rows[col] + hills[row-col+2*n] + dales[row+col]
	return res == 0
}

func backtrack(row, count, n int, rows, hills, dales []int) int {
	for col := 0; col < n; col++ {
		if isNotUnderAttack(row, col, n, rows, hills, dales) {
			rows[col] = 1
			hills[row-col+2*n] = 1
			dales[row+col] = 1

			if row+1 == n {
				count++
			} else {
				count = backtrack(row+1, count, n, rows, hills, dales)
			}

			rows[col] = 0
			hills[row-col+2*n] = 0
			dales[row+col] = 0
		}
	}
	return count
}

func totalNQueens(n int) int {
	rows := make([]int, n)
	hills := make([]int, 4*n-1)
	dales := make([]int, 2*n-1)
	return backtrack(0, 0, n, rows, hills, dales)
}
