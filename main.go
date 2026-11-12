package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	httpStruct := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := httpStruct.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
