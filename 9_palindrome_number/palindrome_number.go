package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	oldX := x
	y := int(0)
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}

	if oldX == y {
		return true
	} else {
		return false
	}
}

func main() {
	x := 121
	is := isPalindrome(x)
	fmt.Println(is)
}
