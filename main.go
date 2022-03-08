package main

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github/aiwantaozi/github-actions-demo/pkg/handler"
)

func main() {
	logrus.Info("Starting run")
	http.HandleFunc("/", handler.HelloHandler)
	http.ListenAndServe(":8080", nil)
}
