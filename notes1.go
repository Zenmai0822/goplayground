package goplayground

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// runes explores the differences between single quotes and double quotes
func runes() {
	fmt.Println("Rune of Dash")
	fmt.Println('-')
	fmt.Println("Actual Dash")
	fmt.Println("-")
}

// typeOfThings explores the type of primitive data
// Note that runes have int32 values
func typeOfThings() {
	fmt.Println(reflect.TypeOf('-')) // int32
	fmt.Println(reflect.TypeOf(684)) // int
	fmt.Println(reflect.TypeOf(3.3)) // float64
	fmt.Println('-' + '_')           // 140
	fmt.Println('-' + '_' + 60)      // 140
}

// makevars explores the grammar for creating variables
func makevars() {
	var zero = 0
	var zerofloat float64
	fmt.Println(zero)
	fmt.Println(zerofloat)
}

// conversions attempts to cast variables to different types
func conversions() {
	length := 3.1
	width := 5
	fmt.Println("Result: (Cast to int)", int(length)*width)
	fmt.Println("Result: (Cast to float64): ", length*float64(width))
	threePointNine := 3.9
	fmt.Println("3.9 as float64: ", threePointNine) // float is dropped entirely
	fmt.Println("3.9 cast to int: ", int(threePointNine))
}

// replacer explores replacer.Replace methods
func replacer() {
	old := "#es#"
	replacer := strings.NewReplacer("#", "t")
	new := replacer.Replace(old)
	fmt.Println(new)
}

// readInput reads input from os.Stdin and prints the type (string) and the input
func readInput() {
	fmt.Print("Input Prompt:\n>")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // returns input and error
	fmt.Println(reflect.TypeOf(input))  // type of input is string (duh)
	fmt.Println(input)
}

// readInputWithLog attempts to read strings from os.Stdin and parse it as integer.
// It contains full error detection

func readInputWithLog() {
	fmt.Print("Input Prompt:\n>")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n') // returns input and error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(input)) // type of input is string (duh)
	fmt.Println(input)
	input = stripSpaces(input)
	inputInt, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(inputInt))
	fmt.Println(inputInt)
}

// scopes explores the scope of variables
func scopes() {
	var funcVar = 1
	if funcVar == 1 {
		// ifBlockVar is only in the if block scope
		var ifBlockVar = 2
		fmt.Println(funcVar, ifBlockVar)
	}
}
