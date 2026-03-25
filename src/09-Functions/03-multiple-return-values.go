/*
Multiple Return Values

Go functions can return multiple values, commonly used for returning result and error. Syntax: func name() (Type1, Type2). Caller receives all returned values or uses blank identifier _ to ignore unwanted values. Idiomatic for error handling pattern.
*/

package main

import "fmt"

func calculateStats(numbers []int) (int, int, float64) {
	if len(numbers) == 0 {
		return 0, 0, 0.0
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	avg := float64(sum) / float64(len(numbers))
	max := numbers[0]

	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return sum, max, avg
}

// Named Return Values
func divideNumbers(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}

	result = a / b
	return
}

func MultipleReturnValues() {
	numbers := []int{10, 20, 30, 40, 50}
	total, maximum, average := calculateStats(numbers)
	fmt.Printf("Total: %d, Maximum: %d, Average: %.2f\n", total, maximum, average)

	result, err := divideNumbers(20, 10)
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}

	fmt.Printf("Division result: %d\n", result)
}
