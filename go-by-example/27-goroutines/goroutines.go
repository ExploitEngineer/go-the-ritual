package main

import (
	"fmt"
	"time"
)

func task(id uint16) {
	fmt.Println("doing task", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		go task(uint16(i))

		func(i uint16) {
			fmt.Println(i)
		}(uint16(i))
	}

	time.Sleep(time.Second * 2)
}
