package somegetters

import "awesomeProject/module3/task1/wordz"

func Digit() string {
	wordz.Words = []string{"One", "Two", "Three", "Four", "Five"}
	return wordz.Random()
}
