package errors

import "fmt"

func PrintErrors() {
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

type CustomErr struct {
	Msg string
}

func (c *CustomErr) Error() string {
	return c.Msg
}

type OtherCustomErr struct {
	Msg string
}

func (c *OtherCustomErr) Error() string {
	return c.Msg
}

func TypeAssertOnCustomError() {
	// Q: How do I type assert on a custom error that implements the error interface
	var e error
	e = &CustomErr{"asdf"}
	_, isCustomErr := e.(*CustomErr)
	fmt.Println(isCustomErr) // true
	_, isOtherCustomErr := e.(*OtherCustomErr)
	fmt.Println(isOtherCustomErr) // false
	// A: Just use regular type assertions
}
