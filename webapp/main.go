package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "loger: ", log.Lshortfile)
	)
	logger.Print("Starting...")
	http.Handle("/", &myHandler{"hi"})
	http.ListenAndServe(":8000", nil)

	fmt.Print(&buf)
}
