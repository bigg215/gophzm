package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllZips(context.Background())

	if err != nil {
		return fmt.Errorf("error reseting database: %w", err)
	}

	return nil
}

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

func (s *state) apiHandlerReady(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK\n")
}

func (s *state) apiHanderLookupZip(w http.ResponseWriter, req *http.Request) {
	zipcode := req.PathValue("zipcode")

	if len(zipcode) == 0 {
		respondwithError(w, http.StatusInternalServerError, "invalid zipcode", nil)
		return
	}

	zoneData, err := s.db.GetZipZone(req.Context(), zipcode)

	if err != nil {
		respondwithError(w, http.StatusNotFound, "zipcode not found", err)
		return
	}
	respondWithJson(w, http.StatusOK, zoneData)
}
