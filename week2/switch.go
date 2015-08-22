package main

import "fmt"

func main() {
	switch {
	default:
		fmt.Println("I am default")
	case false:
		fmt.Println("I am false")
	case true:
		fmt.Println("I am true")
	}
}

// This will print `I am true` because no expression
// is passed to switch statement. So it will evaluate
// the case statement that evaluates to true.
