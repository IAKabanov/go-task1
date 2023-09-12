package task1

import (
	"fmt"
	"time"
)

func Task() {
	printTaskNum()
	fmt.Printf("Today is %s", time.Now().Format("02.01.2006 15:04"))
}

func printTaskNum() {
	fmt.Println()
	fmt.Println()
	fmt.Println("====================Task 1====================")
	fmt.Println()
	fmt.Println()
}
