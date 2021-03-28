package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(str []string) string {
	eleNum := len(str)
	minLen := math.MaxInt64

	for _, v := range str {
		n := len(v)
		if n < minLen {
			minLen = n
		}
	}

	var commStr string

	for i := 0; i < minLen; i++ {
		singStr := str[0][i]
		isComm := true
		for j := 1; j < eleNum; j++ {
			if singStr != str[j][i] {
				isComm = false
				break
			}
		}

		if isComm {
			commStr += string(singStr)
		} else {
			break
		}
	}

	return commStr
}

func main() {
	str := []string{"flower", "flow", "flight"}

	commStr := longestCommonPrefix(str)
	fmt.Println(commStr)
}
