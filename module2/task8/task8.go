package task8

import (
	"awesomeProject/utility"
	"bufio"
	"fmt"
	"os"
)

func Task() {
	utility.AnnounceNewTask(2, 8)

	var pathIn = "./data/08_task_in.txt"
	var fileIn, err = os.Open(pathIn)
	if err != nil {
		fmt.Printf("Cannot open the file %s: cause: %s\n", pathIn, fmt.Errorf(err.Error()))
	} else {
		defer func() {
			var err = fileIn.Close()
			if err != nil {
				fmt.Printf("Cannot close the file %s: cause: %s\n", pathIn, fmt.Errorf(err.Error()))
			}
		}()
		var scanner = bufio.NewScanner(fileIn)
		var lineNum = 0
		for scanner.Scan() {
			lineNum++
		}
		err = scanner.Err()
		if err != nil {
			fmt.Printf("Something went wrong to the scanner, that was scanning to %s. Happened %s\n", pathIn, fmt.Errorf(err.Error()))
		}
		fmt.Printf("Scanner, that was scanning to %s reached EOF\n", pathIn)
		fmt.Printf("Total strings: %d.", lineNum)

	}

}
