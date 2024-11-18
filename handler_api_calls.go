package main

import (
	"io"
	"net/http"
)

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
