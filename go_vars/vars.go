package vars

import "fmt"

func Run(option string) {
	switch (option) {
	case "string":
		runString()
	case "int":
		runInt()
	default:
		fmt.Println("For this option you need to send an option like string, int or...")
	}
}