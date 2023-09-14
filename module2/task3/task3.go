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
	copy(weekdays, days[1:])
	fmt.Printf("Weekdays are %s\n", strings.Join(weekdays, ", "))

	days = append(days[6:], days[:1]...)
	fmt.Printf("Weekends are %s\n", strings.Join(days, ", "))

	allDays := append(weekdays, days...)
	fmt.Printf("All days are %s\n", strings.Join(allDays, ", "))
}
