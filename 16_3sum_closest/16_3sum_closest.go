package main

import (
	"fmt"
	"math"
	"sort"
)

func WithBranch(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	n := len(nums)
	sort.Ints(nums)

	fmt.Println(nums)

	res := math.MaxInt64
	for i := 0; i < n-2; i++ {
		a := nums[i]

		lo := i + 1
		hi := n - 1
		for lo < hi {
			threeSum := a + nums[lo] + nums[hi]
			if threeSum == target {
				return target
			} else if threeSum > target {
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}
				hi--
			} else if threeSum < target {
				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}
				lo++
			}

			value1 := WithBranch(res - target)
			if value1 < 0 {
				value1 = res
			}
			value2 := WithBranch(threeSum - target)
			if value1 > value2 {
				res = threeSum
			}
		}

		for i < n-2 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}

func main() {
	//nums := []int{-1, 2, 1, -4}
	nums := []int{-3, -2, -5, 3, -4}
	target := -1
	fmt.Println(threeSumClosest(nums, target))
}
