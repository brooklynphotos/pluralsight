package main

import (
	"fmt"
	"strings"
)

func main() {
	phrasePrinter()
	selectPrint()
}

func selectPrint() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch1 <- "hello"
	select {
	case x := <-ch1:
		fmt.Printf("got from channel 1: %s\n", x)
	case x := <-ch2:
		fmt.Printf("got from channel 2: %s\n", x)
	default:
		fmt.Println("Got nothing")
	}
}

func phrasePrinter() {
	phrase := "These are those times\n"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	close(ch)

	for w := range ch {
		fmt.Print(w + " ")
	}

}
