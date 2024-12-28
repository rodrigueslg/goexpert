package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fileServer)

	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Content loaded from string"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
