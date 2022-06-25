// package luhn
package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	var digits []int
	var chars []string = strings.Split(strings.ReplaceAll(id, " ", ""), "")

	if len(chars) <= 1 {
		return false
	}

	for _, c := range chars {
		v, err := strconv.Atoi(c)

		if err != nil {
			return false
		}

		digits = append(digits, v)
	}
	var d, result int
	for i := len(digits) - 1; i >= 0; i = i - 1 {

		tmp := digits[i]
		if d%2 == 1 {
			tmp *= 2
			if tmp > 9 {
				tmp -= 9
			}
		}
		result += tmp
		d++
	}
	return (result%10 == 0)
}

// 1224 4454 6745 3124
// 1234 5678 9

// 1 2 2 4 4 4 5 4 6 7  4  5  3  1  2  4
// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15

// i = 16 - 2 => 14
// i = 14 - 2 = 12
// i = 12 - 2 = 10
// ...
// i = 2 - 2 = 0

// i = 16 - 1 = 15

// 059

// 0  5  9
// 0  1  2

// i = 3 - 1 = 2
// i = 2 - 1 = 1
// i = 1 - 1 = 0
