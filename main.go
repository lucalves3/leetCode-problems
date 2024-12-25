package main

import (
	"app/problems"
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15, 30, 21, 123, 55, 10, 32}
	target := 178

	result := problems.Twosum(nums, target)
	fmt.Printf("result: %v\n", result)
}
