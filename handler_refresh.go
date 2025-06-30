package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/ksk/httpserver/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
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

	userWithToken, err := cfg.dbQueries.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid refresh token", err)
		return
	}

	if userWithToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "Refresh token has been revoked", nil)
		return
	}

	if userWithToken.ExpiresAt.Valid && userWithToken.ExpiresAt.Time.Before(time.Now()) {
		respondWithError(w, http.StatusUnauthorized, "Refresh token has expired", nil)
		return
	}

	newAccessToken, err := auth.MakeJWT(userWithToken.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create access token", err)
		return
	}

	type refreshResponse struct {
		Token string `json:"token"`
	}

	response := refreshResponse{
		Token: newAccessToken,
	}

	respondWithJSON(w, http.StatusOK, response)
}
