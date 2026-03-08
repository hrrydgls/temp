package pointer

import (
	"fmt"
	"reflect"
)

func Run () {
	a := 3


	b := &a 

	*b = 4

	fmt.Printf("The value of a is %d\n", a)

	typeOfA := reflect.TypeOf(a)
	typeOfB := reflect.TypeOf(b)

	fmt.Printf("Type of a and b respectivly are: %s, %s\n", typeOfA, typeOfB)



	// pass by refernece

	d := 4

	changeItToFive(&d)

	fmt.Printf("After call by refernece: %d\n", d)

}

func changeItToFive(d *int) {
	*d = *d + 1
}