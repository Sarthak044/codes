package main

import (
	"fmt"
)

func main() {
	var N int64
	var i int64
	var s int64
	fmt.Scan(&N)
	arr := make([]int64, N)
	for i = 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&s)

	found := false
	for i = 0; i < N; i++ {
		if arr[i] == s {
			fmt.Println(i)
			found = true
			break
		}
	}
	if !found {
		fmt.Println(-1)
	}
}