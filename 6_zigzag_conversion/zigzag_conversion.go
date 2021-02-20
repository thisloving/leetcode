package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows <= 1 || len(s) <= numRows {
		return s
	}

	lineSpace := (numRows - 1) * 2
	var conStr string

	str := make([][]string, numRows, numRows)
	for i := 0; i < numRows; i++ {
		str[i] = make([]string, 0, len(s)/numRows+1)
		str[i] = append(str[i], string(s[i]))

		add := 0
		index := 0
		for j := i; j < len(s); {
			for m := 0; m < 1; m++ {
				if add == 0 {
					break
				}
				str[i] = append(str[i], string(s[j]))
			}

			if index%2 == 0 {
				add = lineSpace - (i*2)%lineSpace
			} else {
				add = (i * 2) % lineSpace
			}

			j += add
			index++
		}

		for n := 0; n < len(str[i]); n++ {
			conStr += string(str[i][n])
		}
	}

	return conStr
}

func main() {
	s := "PAYPALISHIRING"
	conStr := convert(s, 4)
	fmt.Println(conStr)
}
