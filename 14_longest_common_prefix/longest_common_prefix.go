package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	eleNum := len(strs)
	minLen := math.MaxInt64

	for _, v := range strs {
		n := len(v)
		if n < minLen {
			minLen = n
		}
	}

	if minLen == 0 || eleNum == 0 {
		return ""
	}

	var commStr string

	for i := 0; i < minLen; i++ {
		singStr := strs[0][i]
		isComm := true
		for j := 1; j < eleNum; j++ {
			if singStr != strs[j][i] {
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
	//str := []string{"flower", "flow", "flight"}
	str := []string{}

	commStr := longestCommonPrefix(str)
	fmt.Println(commStr)
}
