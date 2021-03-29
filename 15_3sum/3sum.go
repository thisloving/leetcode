package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, start int, target int) [][]int {
	lo := start
	hi := len(nums) - 1

	res := make([][]int, 0)
	for lo < hi {
		left := nums[lo]
		right := nums[hi]
		sum := left + right
		if sum < target {
			for lo < hi && left == nums[lo] {
				lo++
			}
		} else if sum > target {
			for lo < hi && right == nums[hi] {
				hi--
			}
		} else {
			pair := make([]int, 0)
			pair = append(pair, left)
			pair = append(pair, right)
			res = append(res, pair)

			for lo < hi && left == nums[lo] {
				lo++
			}

			for lo < hi && right == nums[hi] {
				hi--
			}
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	target := 0

	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		tuples := twoSum(nums, i+1, target-nums[i])
		if len(tuples) == 0 {
			continue
		}

		for _, data := range tuples {
			pair := make([]int, 0)
			for _, v := range data {
				pair = append(pair, v)
			}
			pair = append(pair, nums[i])
			res = append(res, pair)
		}

		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	for _, data := range res {
		for _, v := range data {
			fmt.Printf("%2d ", v)
		}
		fmt.Println()
	}
}
