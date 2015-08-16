package main

import "fmt"

func main() {
	// Formatting booleans
	fmt.Printf("%v\n", false)
	fmt.Printf("%t\n", true)

	// Formatting integers
	fmt.Printf("%v\n", 987)
	fmt.Printf("%d\n", 345)

	// Formatting floats
	fmt.Printf("%v\n", 98753213213.231232131231) // 9.875321321323123e+10
	fmt.Printf("%f\n", 98753213213.231232131231) // 98753213213.231232
	fmt.Printf("%g\n", 98753213213.231232131231) // 9.875321321323123e+10

	// Width and precision
	fmt.Printf("|%6d|%6d|\n", 23, 987)
	fmt.Printf("|%.2f\n", 232132132232.323232)

	// Left justify. Use %-
	fmt.Printf("|%-6d|\n", 23)

	// Sprintf
	s := fmt.Sprintf("%s", "Go Goa Gone")
	fmt.Println(s)
}
