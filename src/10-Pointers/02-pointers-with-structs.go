/*
Pointers with Structs

Pointers to structs enable efficient passing of large structures and allow modification of struct fields. Access fields with (*ptr).field or shorthand ptr.field. Common for method receivers and when structs need to be modified by functions. Essential for memory efficiency.
*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func PointersWithStrcuts() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
