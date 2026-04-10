package variables

import "fmt"

func ShadowingChallange() {
	x := 10

	{
		x := 20
		fmt.Println(x)
	}

	fmt.Println(x)
}
