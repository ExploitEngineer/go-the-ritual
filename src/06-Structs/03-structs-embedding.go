/*
Embedding Structs

Struct embedding includes one struct inside another without field names, making embedded fields directly accessible. Provides composition-based design following Go's philosophy of composition over inheritance. Enables flexible, reusable components.
*/

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base, An embedding looks like a field without a name.
type container struct {
	base
	str string
}

func StructsEmbedding() {
	// When creating s.trcuts with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// we can access the base's fields directly on co, e.g co.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("also num:", co.base.num)

	// Since container embeds base, the methods of base also become methods of a container. Here we invoke a method that was embedded from base directly on co.
	fmt.Println("descrbie:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
