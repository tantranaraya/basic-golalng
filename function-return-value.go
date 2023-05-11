package main

import "fmt"

// function(parameter and typedata)return value typedata{}

func getSomething(title string) string {
	return "My title is " + title
}

func getSay(text string) string {
	if text == "" {
		return "Helo there "
	} else {
		return "Helo " + text
	}
}

func main() {
	result := getSomething("tukang insinyur")
	fmt.Println(result)

	// textResult := getSay("Galaksi 567 alpha proximus")
	// fmt.Println(textResult)

	fmt.Println(getSay("Galaksi 567 alpha proximus"))

}
