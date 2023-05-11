package main

import "fmt"

// Factorial memakai For Loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// Factorial memakai Recursive Function ( function memanggil dirinya sendiri )

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		// value * memanggil function dirinya sendiri
		return value * factorialRecursive(value-1)
	}
}

func main() {

	// Factorial memakai For Loop
	fmt.Println(factorialLoop(5))

	// atau

	loop := factorialLoop(5)
	fmt.Println(loop)

	// dicompare

	fmt.Println(5 * 4 * 3 * 2 * 1)

	// end

	// Factorial memakai Recursive Function
	fmt.Println(factorialRecursive(5))

	// atau

	recursive := factorialRecursive(5)
	fmt.Println("Hasil Recursive Function", recursive)

}
