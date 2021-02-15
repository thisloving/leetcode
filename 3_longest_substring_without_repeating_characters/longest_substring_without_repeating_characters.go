package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	/*
		    if len(s) == 0 {
			    return 0
		    }

			m := make(map[string]int)
			length := int(0)
			left := int(0)
			right := int(0)
			tmpLength := int(0)
			for _, v := range s {
				if pos, ok := m[string(v)]; ok {
					if left < pos+1 {
						left = pos + 1
					}
				}

				m[string(v)] = right
				tmpLength = right - left
				if length < tmpLength {
					length = tmpLength
				}
				right++
			}

			return length + 1
	*/

	if len(s) == 0 {
		return 0
	}

	var bitset [256]int
	for i := 0; i < 256; i++ {
		bitset[i] = -1
	}
	length := int(0)
	left := int(0)
	tmpLength := int(0)
	for right := 0; right < len(s); right++ {
		pos := bitset[s[right]]
		if pos >= 0 && left < pos+1 {
			left = pos + 1
		}

		bitset[s[right]] = right
		tmpLength = right - left
		if length < tmpLength {
			length = tmpLength
		}
	}

	return length + 1
}

func main() {
	s := "pwwkew"
	n := lengthOfLongestSubstring(s)
	fmt.Println(n)
}
