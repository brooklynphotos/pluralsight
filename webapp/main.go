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
	w.Write([]byte(fmt.Sprintf("%v world for %s", mh.greeting, r.URL.Path)))
}

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "loger: ", log.Lshortfile)
	)
	logger.Print("Starting...")
	http.Handle("/interface", &myHandler{"hi"})
	http.HandleFunc("/func", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("func")))
	})
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"a\":1}"))
		w.Header().Add("Content-Type", "application/json")
	})
	http.ListenAndServe(":8000", nil)

	fmt.Print(&buf)
}
