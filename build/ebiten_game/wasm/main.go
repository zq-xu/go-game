package main

import (
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":9080", http.FileServer(http.Dir("."))); err != nil {
		log.Fatal(err)
	}
}
