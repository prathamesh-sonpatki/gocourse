package main

import "fmt"

func main() {
	sum := 0

	// for init; condition; post
	for i := 0; i < 100; i++ {
		sum += i
	}

	fmt.Println(sum)

	sum = 1

	// for condition
	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)

	sum = 1
	// infinite loop with break
	for {
		sum += sum
		fmt.Println(sum)
		if sum > 100 {
			break
		}
	}

L:
	for {
		for {
			fmt.Println("inner loop")
			break L
		}

		fmt.Println("outer loop")
		// break

	}
	fmt.Println("Break from all loops")
}
