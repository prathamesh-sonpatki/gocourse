package main

import "fmt"

func main() {
	i := 0
label:
	fmt.Println(i)
	i++
	if i > 50 {
		goto Escape
	}

	lol(i)
	goto label
Escape:
	fmt.Println("Get out of here...")
}

func lol(i int) {
	j := i + 1
	fmt.Println(j)
}
