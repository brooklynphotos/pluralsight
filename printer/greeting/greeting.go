package greeting

import "fmt"

// alias for a function that takes a string but doesn't return anything
type Printer func(string)

type Salutation struct {
	Name     string
	Greeting string
}

func Greet(sal Salutation, printer Printer) {
	database := map[string]string{
		"Bob": "Mr",
		"Joe": "Dr",
	}
	if _, exists := database[sal.Name]; exists {
		printer(CreateMessage(sal))
	} else {
		fmt.Println("Not there")
	}
}

func CreateMessage(sal Salutation) string {
	return sal.Name + " " + sal.Greeting
}

func CreatePrintFunction(suffix string) Printer {
	return func(str string) {
		fmt.Println(str + suffix)
	}
}
