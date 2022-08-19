package is_test

import (
	"fmt"

	"blake.io/is"
)

func ExampleTrue() {
	var theZeroValueOfBool bool
	if is.True(!theZeroValueOfBool) {
		fmt.Println("theZeroValueOfBool is, in fact, the zero value of bool")
	}

	if is.False(theZeroValueOfBool) {
		fmt.Println("theZeroValueOfBool is the opposite of not the zero value of bool")
	}

	// Output:
	// theZeroValueOfBool is, in fact, the zero value of bool
	// theZeroValueOfBool is the opposite of not the zero value of bool
}
