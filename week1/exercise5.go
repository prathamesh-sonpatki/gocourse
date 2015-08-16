package main

import "fmt"

func main() {
	var f, c float64

	fmt.Println("Enter temperature(In Fahrenheit):")
	fmt.Scanf("%f", &f)

	c = (f - 32.0) * 5.0 / 9.0

	fmt.Printf("Farenheit - %.2f\n", f)
	fmt.Printf("Celcius - %.2f\n", c)
}
