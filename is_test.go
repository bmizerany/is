package is_test

import (
	"fmt"

	"blake.io/is"
)

func ExampleTrue() {
	var theZeroValueOfBool bool
	if is.True(is.False(theZeroValueOfBool)) {
		fmt.Println("theZeroValueOfBool is, in fact, the zero value of bool")
	}

	// Output:
	// theZeroValueOfBool is, in fact, the zero value of bool
}
