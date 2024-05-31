package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "Hello World!\n"); err != nil {
			log.Fatal(err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
