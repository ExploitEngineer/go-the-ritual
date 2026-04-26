package conditionals

import "fmt"

func FallthroughTest() {
	day := "Saturday"

	// With fallthrough, Go forces execution of the next case block, even if next case condition is false.
	switch day {
	case "Suaturday":
		fmt.Println("It's Saturday")
		fallthrough
	case "Sunday":
		fmt.Println("It's Weekend")
	default:
		fmt.Println("Weekday")
	}
}
