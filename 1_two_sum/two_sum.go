package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if pos, ok := m[another]; ok {
			return []int{pos, i}
		}
		m[nums[i]] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	pos := twoSum(nums, target)
	fmt.Println(pos)
}
