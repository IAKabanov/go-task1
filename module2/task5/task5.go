package task5

import (
	"awesomeProject/utility"
	"fmt"
	"sort"
)

func Task() {
	utility.AnnounceNewTask(2, 5)

	var contains = func(a []string, x string) bool {
		for _, word := range a {
			if word == x {
				return true
			}
		}
		return false
	}

	var sliceOfStrings = []string{"integer", "boolean", "float", "string"}
	var comparableTrueString = "string"
	var containsTrue = contains(sliceOfStrings, comparableTrueString)
	fmt.Printf("Contains worked %t\n", containsTrue)
	var fasterContainsTrue = fasterContains(sliceOfStrings, comparableTrueString)
	fmt.Printf("FasterContains worked %t\n", fasterContainsTrue)

	var comparableFalseString = "not existed string"
	var containsFalse = contains(sliceOfStrings, comparableFalseString)
	fmt.Printf("Contains worked %t\n", containsFalse)
	var fasterContainsFalse = fasterContains(sliceOfStrings, comparableFalseString)
	fmt.Printf("FasterContains worked %t\n", fasterContainsFalse)

	var max = getMax(5, 10, 25, 66, 197, 903, 286, 52, 590, 669, 152, 8, 910, 392, 64, 31, 705, 557, 98, 18)

	fmt.Printf("Maximum found %d\n", max)
}

func fasterContains(a []string, x string) bool {
	var dict = make(map[string]byte)
	for _, word := range a {
		dict[word] = 0
	}
	_, isExist := dict[x]
	return isExist
}

func getMax(num ...int) int {
	sort.Slice(num, func(i, j int) bool {
		return num[i] > num[j]
	})
	return num[0]
}
