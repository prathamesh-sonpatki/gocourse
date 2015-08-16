package main

import "fmt"

func main() {
	const secret = 42

	const typedSecret int64 = 42

	fmt.Printf("%T %v\n", secret, secret)
	fmt.Printf("%T %v\n", typedSecret, typedSecret)
}
