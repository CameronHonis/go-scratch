package main

import "fmt"

type MyStruct struct {
	Values [5]int
}

func copy() {
	original := MyStruct{Values: [5]int{1, 2, 3, 4, 5}}

	copy := original // Shallow copy

	copy.Values[0] = 10

	fmt.Println(original.Values[0]) // Output: 1
	fmt.Println(copy.Values[0])     // Output: 10
}
