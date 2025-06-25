package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, r *http.Request) {
	type email struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	userEmail := email{}
	err := decoder.Decode(&userEmail)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	nullString := sql.NullString{
		String: userEmail.Email,
		Valid:  true,
	}

	user, err := cfg.dbQueries.CreateUser(r.Context(), nullString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	responseUser := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
		Email:     user.Email.String,
	}

	jsonResponse, err := json.Marshal(responseUser)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't marshal user to JSON", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}
