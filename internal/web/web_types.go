package web

import (
	"time"

	"github.com/MudassirDev/barter/db/database"
)

type apiConfig struct {
	DB           *database.Queries
	JWTSecretKey string
	ExpiresIn    time.Duration
}

type ResponseUser struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	ID        any    `json:"id"`
}
