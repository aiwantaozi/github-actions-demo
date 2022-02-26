package handler

import (
	"fmt"
	"net/http"
	"log"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("received request", r.URL.String())

	title := "Hello pipeline test, time: "
	date := time.Now().String()
	fmt.Fprintf(w, title+date+"\n")
}