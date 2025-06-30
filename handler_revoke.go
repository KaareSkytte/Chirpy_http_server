package main

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/ksk/httpserver/internal/database"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		respondWithError(w, http.StatusUnauthorized, "Missing authorization header", nil)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		respondWithError(w, http.StatusUnauthorized, "Invalid authorization header format", nil)
		return
	}

	refreshToken := headerParts[1]

	now := time.Now()
	err := cfg.dbQueries.RevokeRefreshToken(r.Context(), database.RevokeRefreshTokenParams{
		Token:     refreshToken,
		RevokedAt: sql.NullTime{Time: now, Valid: true},
		UpdatedAt: sql.NullTime{Time: now, Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke access", nil)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
