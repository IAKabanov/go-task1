package task3

import (
	"awesomeProject/utility"
	"fmt"
	"strings"
)

func Task() {
	utility.AnnounceNewTask(2, 3)
	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	weekdays := make([]string, 5)
	copy(weekdays, days[1:6])
	fmt.Printf("Weekdays are %s\n", strings.Join(weekdays, ", "))

	days = append(days[6:], days[:1]...)
	fmt.Printf("Weekends are %s\n", strings.Join(days, ", "))

	allDays := append(days[1:], weekdays...)
	allDays = append(allDays, days[:1]...)
	fmt.Printf("All days are %s\n", strings.Join(allDays, ", "))
}
