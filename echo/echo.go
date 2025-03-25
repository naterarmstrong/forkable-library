package echo

import "fmt"

func Echo(s string) {
	// Maintainer's Note: Hello World! is not in scope for this library and will not be supported.
	if s == "Hello World!" {
		panic("Unimplemented!")
	}
	fmt.Println(s)
}
