package main

import "fmt"

type MyStruct struct {
	Values [5]int
}

func Copy() {
	copyMap()
}

func copySlice() {
	// Q: does slice re-assignment default to at least a shallow copy?
	original := MyStruct{Values: [5]int{1, 2, 3, 4, 5}}

	copy := original // Shallow copySlice

	copy.Values[0] = 10

	fmt.Println(original.Values[0]) // Output: 1
	fmt.Println(copy.Values[0])     // Output: 10
	// A: yes
}

func copyMap() {
	// Q: does map re-assignment default to at least a shallow copy of the map?
	orig := map[string]int{
		"asdf": 12,
	}
	cp := orig

	cp["asdf"] = 11
	fmt.Println(orig["asdf"]) // 11
	fmt.Println(cp["asdf"])   // 11
	// A: no, both maps share the same data
}
