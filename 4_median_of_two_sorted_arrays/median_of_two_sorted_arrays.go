package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenNum1st := len(nums1)
	lenNum2nd := len(nums2)
	if lenNum1st == 0 && lenNum2nd == 0 {
		return float64(0)
	}

	num1st, num2nd := 0, 0
	if lenNum1st > 0 {
		num2nd = nums1[0]
	}

	if lenNum2nd > 0 && num2nd > nums2[0] {
		num2nd = nums2[0]
	}

	i, j := 0, 0
	length := lenNum1st + lenNum2nd
	mid := length / 2
	index := 0
	temp := num1st
	for i < lenNum1st || j < lenNum2nd {
		if i == lenNum1st {
			temp = nums2[j]
			j++
		} else if j == lenNum2nd {
			temp = nums1[i]
			i++
		} else if i < lenNum1st && j < lenNum2nd {
			if nums1[i] < nums2[j] {
				temp = nums1[i]
				i++
			} else {
				temp = nums2[j]
				j++
			}
		}

		num1st = num2nd
		num2nd = temp
		index++

		if mid+1 == index {
			if length%2 != 0 {
				// odd
				num1st = num2nd
			}
			break
		}
	}

	return float64(num1st+num2nd) / float64(2)
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println(median)
}
