package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MudassirDev/barter/db/database"
	"github.com/google/uuid"
)

func (c *apiConfig) handleCreateSkill() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type Request struct {
			Title string `json:"title"`
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

		userID := r.Context().Value(AUTH_KEY)

		skill, err := c.DB.CreateSkill(context.Background(), database.CreateSkillParams{
			ID:        uuid.New(),
			Title:     req.Title,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		if err != nil {
			if !strings.Contains(err.Error(), "UNIQUE") {
				respondWithError(w, http.StatusInternalServerError, "failed to create a skill", err)
				return
			}
			dbSkill, err := c.DB.GetSkillByTitle(context.Background(), req.Title)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, "failed to create a skill", err)
				return
			}
			skill = dbSkill
		}

		_, err = c.DB.CreateUserSkill(context.Background(), database.CreateUserSkillParams{
			UserID:  userID,
			SkillID: skill.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				respondWithError(w, http.StatusBadRequest, "this skill already exists", err)
				return
			}
			respondWithError(w, http.StatusInternalServerError, "failed to create a skill", err)
			return
		}

		respondWithJSON(w, http.StatusCreated, skill)
	})
}
