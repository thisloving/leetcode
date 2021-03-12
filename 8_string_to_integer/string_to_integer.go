package main

import (
	"fmt"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	negative := 1

	if s[0] == '-' {
		negative = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	tempInt := 0
	for i := 0; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			break
		}

		tempInt = tempInt*10 + int(s[i]-'0')

		if tempInt > (2<<30)-1 {
			break
		}
	}

	tempInt = negative * tempInt
	if tempInt < -(2 << 30) {
		tempInt = -(2 << 30)
	}

	if tempInt > 2<<30-1 {
		tempInt = 2<<30 - 1
	}

	return tempInt
}

func main() {
	s := "-9223372036854775808"
	tempInt := myAtoi(s)
	fmt.Println(tempInt)
}
