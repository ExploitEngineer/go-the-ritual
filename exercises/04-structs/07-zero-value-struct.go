package structs

import "fmt"

func ZeroValueStruct() {
	var user User

	// what are default values and why?
	// default or zero values are nil and can be according to the data types when there are no value in ints the zero value is 0 and in string its empty string or for pointer reference data types its nil
	fmt.Printf("%+v\n", user)
}
