package main

import "fmt"

func main() {

	titles := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	slice := titles[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
}
