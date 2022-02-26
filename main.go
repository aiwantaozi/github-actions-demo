package main

import (
	"net/http"
	"log"

	"github/aiwantaozi/github-actions-demo/pkg/handler"
)

func main() {
	log.Println("Starting run")
	http.HandleFunc("/", handler.HelloHandler)
	http.ListenAndServe(":8080", nil)
}

