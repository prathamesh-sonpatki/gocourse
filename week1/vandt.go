package main

import "fmt"

var x int

// foo := "foo" // Error - ./vandt.go:7: non-declaration statement outside function body

func main() {
	p, q := 10, 20

	type MyInt int
	var r MyInt = 1000

	fmt.Println("p is:", p)
	fmt.Println("q is:", q)
	fmt.Println("r is:", r)

	fmt.Println("x is:", x)

	fmt.Printf("%T %s", "Hello", "Hello\n")

	fmt.Printf("%T %v\n\n", 'x', 'x')

	type MyWeirdCompex complex128
	const I MyWeirdCompex = 0 + 0i
	fmt.Println(I)
}
