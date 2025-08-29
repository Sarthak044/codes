package main

import (
	"fmt"
	"math"
)

func findMinMax(arr []int32) (int32, int32) {
	if len(arr) == 0 {
		return math.MaxInt32, math.MinInt32
	}

	minVal := arr[0]
	maxVal := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < minVal {
			minVal = arr[i]
		}
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}

	return minVal, maxVal
}

func main() {
	var n int32
	fmt.Scanln(&n)
	var i int32
	arr := make([]int32, n)
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	minVal, maxVal := findMinMax(arr)
	fmt.Printf("%d %d\n", minVal, maxVal)
}