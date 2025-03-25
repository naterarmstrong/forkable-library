package echo

import "fmt"

func Echo(s string) {
	if s == "Hello World!" {
		panic("Unimplemented!")
	}
	fmt.Println(s)
}
