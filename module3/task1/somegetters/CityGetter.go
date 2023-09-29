package somegetters

import "awesomeProject/module3/task1/wordz"

func City() string {
	wordz.Words = []string{"Purrville", "Clawtown", "Meowcago", "New Furrbelly", "Los Kittios"}
	return wordz.Random()
}
