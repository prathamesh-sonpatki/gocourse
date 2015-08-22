package main

import "fmt"

func main() {
	i := 5
	switch i {
	case 4:
		fmt.Println("number 4")
	case i > 8:
		fmt.Println("i is > 8")
	default:
		fmt.Println("default")
	}

}

/*
 This will fail to compile because of type mismatch
 in second case expression where that expression
 evaluates to bool instead of int which is the type
 of switch statement.
*/
