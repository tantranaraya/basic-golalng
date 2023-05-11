package main

import "fmt"

func main() {
	var hasil = 12 * 19

	var a = 3
	var b = 5
	var c = a + b
	var d = a * b

	fmt.Println("hasil adalah", hasil)
	fmt.Println(c)
	fmt.Println(d)

	// Augmented Assignments
	// a += 10 -> a = a + 10
	// a -= 10 -> a = a - 10
	// a *= 10 -> a = a * 10
	// a /= 10 -> a = a / 10
	// a %= 10 -> a = a % 10

	var e = 10
	e += 13
	fmt.Println(e)

	// Unary Operator
	//  ++ -> a = a + 1
	//  -- -> a = a - 1
	//  - Negative
	//  + Positive
	//  ! Boolean kebalikan

	var f = 12
	f++
	fmt.Println(f)
}
