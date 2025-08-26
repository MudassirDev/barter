package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MudassirDev/barter/db/database"
)

func (c *apiConfig) handleSearch(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		City string `json:"title"`
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

	result, err := c.DB.GetUsersWithCity(context.Background(), formatQuery(req.City))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to get users", err)
		return
	}

	var skills []database.GetSkillsByUserIDRow

	for _, user := range result {
		skill, _ := c.DB.GetSkillsByUserID(context.Background(), user.ID)
		skills = append(skills, skill...)
	}

	respondWithJSON(w, http.StatusOK, skills)
}

func formatQuery(s string) string {
	return "%" + s + "%"
}
