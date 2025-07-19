package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/sentences", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sentence service")
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SENTENCE_REST_PORT"), mux))
}
