package goplayground

import (
	"fmt"
)

// MainPartThree runs functions in this file, change it as needed.
func MainPartThree() {
	printPointers()
	mutatePointers()
}

// passByValueDemo demonstrates that go copies variable parameters, instead of passing the reference of the parameter variables.
func passByValueDemo() {
	three := 3
	fmt.Printf("Value of variable 'three' is %d\n", three)
	fmt.Printf("Value of function 'plusOne' return value is %d\n", plusOne(three))
	fmt.Printf("Value of variable 'three' is %d\n", three)
}

// plusOne adds one to the given int parameter, and returns it.
// plusOne is called by passByValueDemo() to demonstrate that go is pass by value.
func plusOne(num int) int {
	return num + 1
}

// printPointers prints a variable, the pointer of the variable, and the go representation ("%#v") of the variable.
func printPointers() {
	three := 3
	pointerOfThree := &three
	fmt.Printf("Value of variable 'three': %d\n", three)
	fmt.Printf("Value of variable 'pointerOfThree': %v\n", pointerOfThree)
	fmt.Printf("Literal Value of variable 'pointerOfThree', printed using %%#v: %#v\n", pointerOfThree)
}

func mutatePointers() {
	three := 3
	pointerOfThree := &three
	fmt.Printf("Value of variable 'three': %d\n", three)
	plusOnePointer(pointerOfThree)
	fmt.Printf("Value of variable 'three', after mutation by pointer: %d\n", three)
}

func plusOnePointer(numPointer *int) {
	*numPointer++
}
