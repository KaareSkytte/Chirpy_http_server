package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ksk/httpserver/internal/auth"
)

func (cfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	api_key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	if api_key != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Wrong API Key", nil)
		return
	}

	type requestWebhook struct {
		Event string `json:"event"`
		Data  struct {
			UserID string `json:"user_id"`
		} `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	request := requestWebhook{}
	err = decoder.Decode(&request)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if request.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	userUUID, err := uuid.Parse(request.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "invalid user ID", err)
		return
	}

	_, err = cfg.dbQueries.UpgradeUser(r.Context(), userUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't upgrade user", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
