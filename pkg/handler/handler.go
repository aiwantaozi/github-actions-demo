package handler

import (
	"fmt"
	"net/http"
	"github.com/sirupsen/logrus"
	"time"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("received request", r.URL.String())

	title := "Hello pipeline test, time: "
	date := time.Now().String()
	fmt.Fprintf(w, title+date+"\n")
}