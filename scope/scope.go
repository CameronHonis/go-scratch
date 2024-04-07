package scope

import "fmt"

func InnerFunctionScope() {
	// Q: can functions write to variables declared in the same scope?
	var someVar int64
	var somePtr *[]int64
	fmt.Println(someVar)
	fmt.Println(somePtr)
	func() {
		someVar = 12
		slice := make([]int64, 0)
		somePtr = &slice
	}()
	fmt.Println(someVar)
	fmt.Println(somePtr)
}
