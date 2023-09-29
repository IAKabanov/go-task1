package task7

import (
	"awesomeProject/utility"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Task() {
	utility.AnnounceNewTask(2, 7)

	var pathIn = "./data/07_task_in.txt"
	var fileIn, _ = os.Open(pathIn)
	defer fileIn.Close()
	var scanner = bufio.NewScanner(fileIn)

	var pathOut = "./data/07_task_out.txt"
	var fileOut, _ = os.Create(pathOut)
	defer fileOut.Close()
	var lineNum = 0

	defer func() {
		if panicErr := recover(); panicErr != nil {
			switch panicErr {
			case "parse error":
				fmt.Printf("parse error: empty field on string %d\n", lineNum)
				fmt.Printf("\nList of written:\n\n")
				var pathOut = "./data/07_task_out.txt"
				var fileOut, _ = os.Open(pathOut)
				var scannerOut = bufio.NewScanner(fileOut)
				for scannerOut.Scan() {
					fmt.Println(scannerOut.Text())
				}
			default:
				panic("critical")
			}
		}
		fileOut.Close()
	}()

	for scanner.Scan() {
		lineNum++
		var name, address, city = parse(scanner.Text())
		var outString = fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", lineNum, name, address, city)
		fileOut.WriteString(outString)
	}

	fmt.Printf("File %s was written. Written lines: %d\n", pathOut, lineNum)
}

func parse(text string) (name string, address string, city string) {
	var splittedText = strings.Split(text, "|")
	if len(splittedText) != 3 {
		panic("parse error")
	}
	name = splittedText[0]
	address = splittedText[1]
	city = splittedText[2]

	return name, address, city
}
