package main

import (
	"fmt"
)

func main() {
	var arr []int32
	arr = []int32{1, 1, 0, -1, -1}
	var n int
	n = len(arr)
	fmt.Printf("len(arr) == : %d\n", n)

	var sumOfPositiveValue float32
	var sumOfNegativeValue float32
	var zeroValue float32

	for i := 0; i < n; i++ {
		fmt.Printf("arr[%d] == : %d\n", i, arr[i])
		if arr[i] > 0 {
			sumOfPositiveValue += 1
		}
		if arr[i] < 0 {
			sumOfNegativeValue += 1
		}
		if arr[i] == 0 {
			zeroValue += 1
		}
	}
	fmt.Printf("sumOfPositiveValue/n == : %f\n", sumOfPositiveValue/float32(n))
	fmt.Printf("sumOfNegativeValue/n == : %f\n", sumOfNegativeValue/float32(n))
	fmt.Printf("zeroValue/n == : %f\n", zeroValue/float32(n))
}
