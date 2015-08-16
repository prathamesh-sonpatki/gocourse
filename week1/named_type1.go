package main

import "fmt"

type Duration int64

func main() {
	var duration Duration
	fmt.Println(duration)

	nanosecond := int64(10)

	duration = nanosecond // Error because of type mismatch
}
