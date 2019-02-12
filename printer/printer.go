package main

import "brooklyn.photos/pluralsight/printer/greeting"

func main() {
	var s = greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!!"))
}
