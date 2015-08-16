package main

import "fmt"

type Duration uint

func main() {
	var duration Duration
	fmt.Println(duration)

	nanosecond := int(-10)

	duration = Duration(nanosecond)

	// How does this type conversion happens from int to uint?
	fmt.Println(duration) // 18446744073709551606
}
