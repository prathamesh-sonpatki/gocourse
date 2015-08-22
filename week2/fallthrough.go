package main

import "fmt"

func main() {
	i := 10

	switch i {
	case 10:
		fmt.Println(" Less than or equal to 10")
		fallthrough

	case 11:
		fmt.Println(" Less than or equal to 11")
	case 12:
		fmt.Println(" Less than or equal to 12")
	}
}
