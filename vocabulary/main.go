package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/vocabularies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Vocabulary service")
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("VOCABULARY_REST_PORT"), mux))
}
