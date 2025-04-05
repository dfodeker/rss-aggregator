package main

import (
	"log"
	"net/http"
)

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
	log.Printf("Received request: %s to %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

}
