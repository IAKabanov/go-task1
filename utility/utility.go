package utility

import (
	"fmt"
	"strconv"
)

func getModuleTaskString(module int, task int) string {
	return "==================== Module " + strconv.Itoa(module) + " Task " + strconv.Itoa(task) + " ===================="
}

func AnnounceNewTask(module int, task int) {
	fmt.Println()
	fmt.Println()
	fmt.Println(getModuleTaskString(module, task))
	fmt.Println()
	fmt.Println()
}
