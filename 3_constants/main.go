package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Constants in Go")
	const PI float32 = 3.14
	fmt.Println("Value of PI is:", PI)
	fmt.Printf("Type of PI is: %T\n", PI)
	// PI = 3.145 // Error: cannot assign to PI (constant)``

	// Untyped constant
	const n = 5000000000
	const d = 3e20 / n
	fmt.Println(n)
	fmt.Println(d)

	// Type conversion
	fmt.Println(int64(d))

	// Math package
	fmt.Println(math.Sin(n))

}
