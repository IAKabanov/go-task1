package task6

import (
	"awesomeProject/utility"
	"bufio"
	"fmt"
	"os"
	"time"
)

func Task() {
	utility.AnnounceNewTask(2, 6)
	defer logTime()
	startTime = time.Now()
	var pathIn = "./data/06_task_in.txt"
	var fileIn, _ = os.Open(pathIn)
	defer fileIn.Close()
	var scanner = bufio.NewScanner(fileIn)

	var pathOut = "./data/06_task_out.txt"
	var fileOut, _ = os.Create(pathOut)
	var i = 1
	var bytes = 0
	defer func() {
		fileOut.Close()
		fmt.Printf("File %s has been written! Wrote %d lines, %d bytes\n", fileOut.Name(), i, bytes)
	}()

	for scanner.Scan() {
		var stringToWrite = fmt.Sprintf("%d %s\n", i, scanner.Text())
		var bts, _ = fileOut.WriteString(stringToWrite)
		bytes += bts
		i++
	}
}

var startTime time.Time

func logTime() {
	fmt.Printf("Reading and writing took %d mks\n", time.Since(startTime).Microseconds())
}
