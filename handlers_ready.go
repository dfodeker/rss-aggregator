package main

import (
	"log"
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received request: %s to %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

	respondWithJSON(w, 200, struct{}{})

}
