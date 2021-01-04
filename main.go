package main

import (
	_ "embed"
	"net/http"
	"log"
)

func main() {
	//go:embed "hello.txt"
	var s string

	log.Fatal(http.ListenAndServe(":8083", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(s))
	})))
}
