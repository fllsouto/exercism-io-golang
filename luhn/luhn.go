// package luhn
package luhn

import (
	"fmt"
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
			fmt.Printf("Error! %s", err.Error())
			return false
		}
	}
	fmt.Printf("B idArray: %v\n", idArray)
	for i := len(idArray) - 2; i >= 0; i = i - 2 {
		//fmt.Printf("[%d] -- %d\n", i, idArray[i])
		tmp := idArray[i] * 2
		if tmp > 9 {
			tmp -= 9
		}

		idArray[i] = tmp
		//fmt.Printf("[%d] -- %d\n", i, idArray[i])
	}
	fmt.Printf("A idArray: %v\n", idArray)
	result := 0
	for _, x := range idArray {
		result += x
	}
	validation := (result%10 == 0)
	fmt.Printf("Result: %d -- Validation %t\n\n", result, validation)
	return validation
}
