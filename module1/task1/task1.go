package task1

import (
	"awesomeProject/utility"
	"fmt"
	"time"
)

func Task() {
	utility.AnnounceNewTask(1, 1)
	fmt.Printf("Today is %s", time.Now().Format("02.01.2006 15:04"))
}
