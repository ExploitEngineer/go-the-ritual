package main

import "fmt"

func plus(a uint8, b uint8) uint8 {
	return a + b
}

func plusPlus(a, b, c uint8) uint8 {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
