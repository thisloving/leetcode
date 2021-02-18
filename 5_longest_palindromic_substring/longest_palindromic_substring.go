package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	str := "#"
	for _, v := range s {
		str += string(v)
		str += "#"
	}

	mirror := 0
	center := 0
	i := 0
	maxRight := 0
	p := make([]int, len(str))

	maxLen := 0
	begin := 0
	end := 0
	for i = 0; i < len(str); i++ {
		if i <= maxRight {
			mirror = 2*center - i
			p[i] = maxRight - i
			//fmt.Println(i, p[i], center, mirror, maxRight)
			if mirror >= 0 && p[mirror] < p[i] {
				p[i] = p[mirror]
			}
		}

		left := i - p[i] - 1
		right := i + p[i] + 1
		//fmt.Println("i", i, "left:", left, "right:", right, p[i], "mirror:", mirror, p[mirror], maxRight)
		for j := 0; right < len(str) && left >= 0 && str[left] == str[right]; j++ {
			p[i]++
			right++
			left--
		}

		if maxRight < i+p[i] {
			maxRight = i + p[i]
			center = i
		}

		if p[i] > maxLen {
			maxLen = p[i]
			begin = i - p[i]
			end = i + p[i]
			//fmt.Println("maxLen:", maxLen, "begin:", begin, i, p[i], maxRight)
		}

		//fmt.Println("i", i, string(str[i]), p[i], maxRight, maxLen)
	}

	//fmt.Println(begin, end, maxLen)
	var subStr string
	//subStr = str[begin : end+1]
	//subStr = strings.Replace(subStr, "#", "", -1)
	for i := begin; i <= end; i++ {
		if i%2 != 0 {
			subStr += string(str[i])
		}
	}
	return subStr
}

func main() {
	s := "cbbdbbcccx"
	subStr := longestPalindrome(s)
	fmt.Println(subStr)
}
