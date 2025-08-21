package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MudassirDev/barter/db/database"
	"github.com/MudassirDev/barter/internal/auth"
	"github.com/google/uuid"
)

func (c *apiConfig) handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	var req Request
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		respondWithError(w, http.StatusBadRequest, "invalid content type, expected: application/json", fmt.Errorf("invalid content type"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&req); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error", err)
		return
	}

	password, err := auth.HashPassword(req.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "password doesn't match the constraints", err)
		return
	}

	user, err := c.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		if !strings.Contains(err.Error(), "UNIQUE") {
			respondWithError(w, http.StatusInternalServerError, "failed to create user", err)
			return
		}
		duplicateKey := strings.Split(err.Error(), "users.")[1]
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Duplicate Key: %v", duplicateKey), err)
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (c *apiConfig) handleLoginUser(w http.ResponseWriter, r *http.Request) {}
