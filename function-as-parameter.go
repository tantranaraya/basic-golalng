package main

import "fmt"

// disini function sebagai parameter -> func sayWordsFilter(param1 string, param2 func(string) string){}

// bentuk pertama
type Filter func(string) string

func sayWordsFilter(word string, filter Filter) {
	wordFiltered := filter(word)
	fmt.Println(wordFiltered)
}

// bentuk kedua
// func sayWordsFilter(word string, filter func(string) string) {
// 	wordFiltered := filter(word)
// 	fmt.Println(wordFiltered)
// }

func spamFilter(word string) string {
	if word == "Fuck" {
		return "mencegat " + word + " ... kata2 tidak dimengerti"
	} else {
		return "mengijinkan " + word
	}
}

func main() {
	// disini juga func spamFilter sebagai parameter dari func sayWordsFilter
	sayWordsFilter("nandini", spamFilter)

	// penggunaan init statement
	filter := spamFilter
	sayWordsFilter("Fuck", filter)
}
