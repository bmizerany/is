package is

import "fmt"

func ExampleTrue() {
	var theZeroValueOfBool bool
	if True(!theZeroValueOfBool) {
		fmt.Println("theZeroValueOfBool is not the zero value of bool")
	}

	if False(theZeroValueOfBool) {
		fmt.Println("theZeroValueOfBool is the opposite of not the zero value of bool")
	}

	// Output:
	// theZeroValueOfBool is not the zero value of bool
	// theZeroValueOfBool is the opposite of not the zero value of bool
}
