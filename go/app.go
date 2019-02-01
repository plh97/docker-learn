package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from the api!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("listening on 80")
	log.Fatal(http.ListenAndServe(":80", nil))
}