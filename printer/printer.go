package main

import "fmt"

// alias for a function that takes a string but doesn't return anything
type Printer func(string)

func Greet(msg string, printer Printer) {
	printer(msg)
}

func CreatePrintFunction(suffix string) Printer {
	return func(str string) {
		fmt.Println(str + suffix)
	}
}

func main() {
	// var s = Salutation("Bob", "Hello")
	var s = "Hi"
	Greet(s, CreatePrintFunction("!!!"))
}
