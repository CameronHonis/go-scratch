package main

import "fmt"

func errors() {
	printErrors()
}

func printErrors() {
	// do errors need to be referenced by calling their `Error()` method?
	fmt.Printf("this is an: %s\n", returnErr().Error())
	fmt.Printf("this is an: %s\n", returnErr())
	// no. these are equivalent

	// does fmt only format errors as string if encapsulated in an f-string?
	fmt.Println("this is an:", returnErr().Error())
	fmt.Println("this is an:", returnErr())
	// no. there are equivalent
}

func returnErr() error {
	return fmt.Errorf("error message")
}
