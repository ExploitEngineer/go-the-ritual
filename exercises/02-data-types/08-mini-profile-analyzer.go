package datatypes

import "fmt"

func MiniProfileAnalyzer() {
	name := "Abdul Rafay"
	age := 18
	height := 5.10
	grade := 'A'
	isStudent := true

	fmt.Println("=====================")
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My age is %d\n", age)
	fmt.Printf("My height is %f\n", height)
	fmt.Printf("My grade is %c\n", grade)
	fmt.Printf("I am a student %t\n", isStudent)
	fmt.Println("=====================")
}
