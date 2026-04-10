package variables

import "fmt"

func IotaEnum() {
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println(Friday)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Saturday, Sunday)
}
