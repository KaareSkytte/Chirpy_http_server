package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.dbQueries.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't fetch chirps", err)
		return
	}

	mainChirps := []Chirp{}

	for _, dbChirp := range dbChirps {
		mainChirps = append(mainChirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		})
	}

	jsonData, err := json.MarshalIndent(mainChirps, "", " ")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't Marshal data", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
