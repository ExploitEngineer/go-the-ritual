package compositetypes

import "fmt"

func SliceAppendGrowth() {
	s := []int{1, 2, 3}

	fmt.Println(len(s))
	fmt.Println("cap: ", cap(s))

	s = append(s, 20)
	fmt.Println(s)

	// When does Go allocate new memory?
	// cause when we first defiend the slice it has the capacity size same as his length and when we append the new element in slice it should get more memory to make the capacity big cause there is no capacity to store more so that's why it allocates new memory.
	// append may or may not allocate depending on remaining capacity
	fmt.Println("cap: ", cap(s))
}
