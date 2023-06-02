package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("got request: ", r.URL.Path)
		w.Write([]byte("Hi!"))
	})
	PORT := "8080"
	log.Println(fmt.Sprintf("listening on port: %v", PORT))
	log.Println(http.ListenAndServe(":"+PORT, mux))
}
