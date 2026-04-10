package introduction

import (
	"fmt"
	// "os" - i am getting imported os but not used its because we are not using this but adding this so compiler is giving the error that you cannot do this because if you are not using this why take more memory.
)

func ImportErrorTest() {
	fmt.Println("hey i am using fmt package")
}
