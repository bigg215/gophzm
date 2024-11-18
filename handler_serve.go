package main

import (
	"log"
	"net/http"
)

func handlerServe(s *state, cmd command) error {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/status", s.apiHandlerReady)
	mux.HandleFunc("GET /api/zip/{zipcode}", s.apiHanderLookupZip)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Starting Server ...")
	log.Fatal(srv.ListenAndServe())
	return nil
}
