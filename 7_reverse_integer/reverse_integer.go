package main

import "fmt"

func reverse(x int) int {
	temp := 0
	for x != 0 {
		temp = temp*10 + x%10
		x = x / 10
		//fmt.Println(x, temp)
	}

	if temp < -(1<<31) || temp > (1<<31)-1 {
		temp = 0
	}

	return temp
}

func main() {
	x := -1000
	y := reverse(x)
	fmt.Println(y)
}
