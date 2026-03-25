/*
Embedding Interfaces

Create new interfaces by combining existing ones, promoting composition and reusability. Embedded interface methods automatically included. Enables interface hierarchies from simpler, focused interfaces. Supports composition over inheritance for modular, extensible systems.
*/

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// A container embeds a base. An embeding looks like a field without a name.
type container struct {
	base
	str string
}

func EmbeddingInterfaces() {
	// when creating struct with literals, we have to initialize the embedding explicitly; here the embedded type serves as the field name.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// we can access the base field directly on co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.num)

	// Since container embeds base, the methods of base also become methods of a container. Here we invoke a method that was embedded from base directly on co.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Embedding structs with methods may be used to bestow interface implementations onto other structs. Here we see that a container now implements the describer interface because it embeds base.
	var d describer = co
	fmt.Println("describe :", d.describe())
}
