package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's Golang in WEB"))
}

func main() {
	name, age := "Ivan", 24
	acquaintance := fmt.Sprintf("I'm %s, and my age %d", name, age)
	println(acquaintance)

	http.HandleFunc("/", Handler)

	log.Println("Start HTTP server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
