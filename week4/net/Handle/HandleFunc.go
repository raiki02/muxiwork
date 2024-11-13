package main

import (
	"io"
	"net/http"
)

func main() {
	HelloFunc1 := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, `HelloFunc1 := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, ???)`)
	}
	HelloFunc2 := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, `HelloFunc2 := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, Haha \n)
	} \n`)
	}

	http.HandleFunc("/", HelloFunc1)
	http.HandleFunc("/HelloFunc2", HelloFunc2)

	port := ":8080"
	http.ListenAndServe(port, nil)
}
