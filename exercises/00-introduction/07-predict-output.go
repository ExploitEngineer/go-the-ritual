package introduction

import "fmt"

func PredictOutput() {
	fmt.Println("Go")
	fmt.Print("Lang")

	/*
				Output:
				Go
				Lang%

		the % is because in the first line we are using Println function which auto adds the \n escape sequence character at the end but the normal Print function does not that's why
	*/
}
