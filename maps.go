package goplayground

import "fmt"

// MapPlayground is a function that plays around map structures of go.
// It does not take in or return anything.
func MapPlayground() {
	// Create the map variable, but currently nil
	var mapOne map[int]string
	fmt.Println("Created variable mapOne, before 'make'")
	fmt.Printf("Value of mapOne: %v\n", mapOne)
	fmt.Printf("Value of mapOne, internal representation: %#v\n", mapOne)
	fmt.Printf("Getting anything from this un-inited map returns nothing: %v\n", mapOne[1])
	fmt.Printf("Getting anything from this un-inited map returns nothing (internal rep.): %#v\n", mapOne[1])
	res, ok := mapOne[1]
	fmt.Printf("var 'res' from getting an un-init map: %v\n", res)
	fmt.Printf("var 'res' from getting an un-init map (internal rep.): %#v\n", res)
	fmt.Printf("var 'ok' from getting an un-init map: %v\n", ok)
	fmt.Printf("var 'ok' from getting an un-init map (internal rep.): %#v\n", ok)
	// Using 'make' to init the map
	mapOne = make(map[int]string)
	fmt.Println("\nInitialized map using make()")
	fmt.Printf("Value of mapOne %v, internal rep. is %#v\n", mapOne, mapOne)
	res, ok = mapOne[1]
	fmt.Printf("var 'res' from getting an init but empty map: %v\n", res)
	fmt.Printf("var 'res' from getting an init but empty map (internal rep.): %#v\n", res)
	fmt.Printf("var 'ok' from getting an init but empty map: %v\n", ok)
	fmt.Printf("var 'ok' from getting an init but empty map (internal rep.): %#v\n", ok)

	fmt.Println("\nFilling data")
	mapOne[1] = "hello"
	res, ok = mapOne[1]
	fmt.Printf("var 'res' from getting a filled map: %v\n", res)
	fmt.Printf("var 'res' from getting a filled map (internal rep.): %#v\n", res)
	fmt.Printf("var 'ok' from getting a filled map: %v\n", ok)
	fmt.Printf("var 'ok' from getting a filled map (internal rep.): %#v\n", ok)
}

/*
Run Result:

Created variable mapOne, before 'make'
Value of mapOne: map[]
Value of mapOne, internal representation: map[int]string(nil)
Getting anything from this un-inited map returns nothing:
Getting anything from this un-inited map returns nothing (internal rep.): ""
var 'res' from getting an un-init map:
var 'res' from getting an un-init map (internal rep.): ""
var 'ok' from getting an un-init map: false
var 'ok' from getting an un-init map (internal rep.): false

Initialized map using make()
Value of mapOne map[], internal rep. is map[int]string{}
var 'res' from getting an init but empty map:
var 'res' from getting an init but empty map (internal rep.): ""
var 'ok' from getting an init but empty map: false
var 'ok' from getting an init but empty map (internal rep.): false

Filling data
var 'res' from getting a filled map: hello
var 'res' from getting a filled map (internal rep.): "hello"
var 'ok' from getting a filled map: true
var 'ok' from getting a filled map (internal rep.): true
*/
