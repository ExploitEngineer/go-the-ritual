package conditionals

import (
	"fmt"
)

func SwitchDayChecker() {
	day := "Sunday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Sunday", "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}
}
