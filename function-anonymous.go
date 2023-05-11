package main

import "fmt"

// anonymous function (func tanpa nama) func(param typedata) return typedata(){}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("mencegat ", name, ".. anda tidak bisa lewat")
	} else {
		fmt.Println("mengijinkan ", name)
	}
}

func main() {

	blacklist := func(name string) bool {
		return name == "fuck"
	}

	registerUser("neira", blacklist)

	// atau

	registerUser("draynerys", func(name string) bool {
		return name == "fuck"
	})
}
