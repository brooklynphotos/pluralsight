package main

import "fmt"

import "brooklyn.photos/pluralsight/printer/greeting"

func main() {
	// start getting data
	routineDone := make(chan string)
	codes := []string{"IN", "CH", "US"}
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
	count := 0
	for s := range routineDone {
		fmt.Println(s)
		count += 1
		if count == len(codes) {
			close(routineDone)
		}
	}
}

func makeNoise(p greeting.Pronounceable) {
	p.Pronounce("hello")
}
