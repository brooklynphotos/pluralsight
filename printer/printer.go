package main

import "fmt"

import "brooklyn.photos/pluralsight/printer/greeting"

func main() {
	// start getting data
	codes := []string{"IN", "CH", "US"}
	routineDone := make(chan string, len(codes))
	for _, code := range codes {
		go func(v string) {
			routineOut := greeting.GetCountry(v)
			routineDone <- routineOut
		}(code)
	}
	s := greeting.Salutation{"Bob", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("!!!"))
	makeNoise(greeting.Child{"Bye", 2})
	mi := greeting.MyInt(2)
	makeNoise(mi)
	for range codes {
		fmt.Println(<-routineDone)
	}
}

func makeNoise(p greeting.Pronounceable) {
	p.Pronounce("hello")
}
