package main

import "fmt"

// function()(multiple return value typedata){}

func getProduct() (string, string) {
	return "krupuk udang", "krupuk reguler"
}

func main() {
	productName, productCategory := getProduct()
	fmt.Println(productName, productCategory)
}
