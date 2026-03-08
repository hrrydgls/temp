package vars

import (
	"fmt"
	"reflect"
)

// lowercase because i dont want to call it from outside
func runString() {

	fmt.Println("Testing string: ")
	
	// var name string

	// name = "Harry"

	name := "Harry"

	fmt.Printf("My name is %s\n", name)

	typeOfString := reflect.TypeOf(name)

	fmt.Printf("Type was %s\n", typeOfString)
}