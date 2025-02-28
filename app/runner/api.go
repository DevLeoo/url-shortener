package runner

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/app/api/router"
	"url-shortener/app/config"
)

type APIRunner struct{}

func (a *APIRunner) Run() {
	port := config.Port
	router := router.Create()
	fmt.Printf("Listening at %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
