package main

import (
	"fmt"
	utf "unicode/utf8"
)

func main() {
	fmt.Printf("Hello")
	fmt.Printf(" World\n")

	print("Learning Go")
	println(".. and enjoying it!")

	s := "Hello, 世界世界"
	fmt.Println(s)
	fmt.Println(utf.RuneCountInString(s))
}
