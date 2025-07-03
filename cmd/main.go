package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/heapsort", postHeapSort)
	log.Println("Starting server on :8080")
	http.ListenAndServe("localhost:8080", router)
}
