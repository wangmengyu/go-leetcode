package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{4, 1, 6, 2, 5, 1, 8, 9}
	res := findMax(arr, 0, len(arr)-1)
	fmt.Println(res)
}
func findMax(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	mid := l + (r-l)>>1
	leftMax := findMax(arr, 0, mid)
	rightMax := findMax(arr, mid+1, r)
	max := int(math.Max(float64(leftMax), float64(rightMax)))
	fmt.Println("max=", max)
	return max
}
