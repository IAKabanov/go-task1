package task4

import (
	"awesomeProject/utility"
	"bytes"
	"fmt"
	"strconv"
)

func Task() {
	utility.AnnounceNewTask(2, 4)
	var library = map[string]map[string][]string{
		"Kogtie": {
			"books":    {"How to be a good cat", "Making humans to clean the toilet easy"},
			"periodic": {"Turkey recipes for cats", "Assembling couch: have a part every month"},
		},
		"Barsik": {
			"books":    {"The crime and the purrishment", "War and purr", "Meowster and Purrgarita"},
			"periodic": {"Hot purring"},
		},
	}

	var plural = ""
	var has = "ve"
	var lenLibrary = len(library)
	if lenLibrary != 1 {
		plural = "s"
		has = "s"
	}
	fmt.Printf("%d customer%s ha%s books or periodic magazines\n", lenLibrary, plural, has)

	for name, books := range library {
		var finalString bytes.Buffer
		finalString.WriteString(name)
		finalString.WriteString(" has")
		var total int
		var hasNext = false
		for readType, names := range books {
			if hasNext {
				finalString.WriteString(" and")

			}
			var lenNames = len(names)
			total += lenNames
			finalString.WriteString(" " + strconv.Itoa(lenNames) + " ")
			finalString.WriteString(readType)
			hasNext = true
		}
		finalString.WriteString(". Total is " + strconv.Itoa(total) + "\n")
		fmt.Print(finalString.String())
	}

}
