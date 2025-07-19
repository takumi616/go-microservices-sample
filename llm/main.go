package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/llms", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "LLM service")
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("LLM_REST_PORT"), mux))
}
