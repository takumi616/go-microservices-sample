package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func setupMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/llms", llmHandler)

	return mux
}

func llmHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "LLM service")
}

func main() {
	log.Fatal(http.ListenAndServe(":"+os.Getenv("LLM_REST_PORT"), setupMux()))
}
