// package luhn
package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	var idArray []int
	var sArray []string = strings.Split(strings.ReplaceAll(id, " ", ""), "")

	if len(sArray) <= 1 {
		return false
	}

	for _, c := range sArray {
		v, err := strconv.Atoi(c)
		idArray = append(idArray, v)

		if err != nil {
			return false
		}
	}
	for i := len(idArray) - 2; i >= 0; i = i - 2 {
		tmp := idArray[i] * 2
		if tmp > 9 {
			tmp -= 9
		}

		idArray[i] = tmp
	}
	result := 0
	for _, x := range idArray {
		result += x
	}
	return (result%10 == 0)
}
