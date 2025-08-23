package web

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/MudassirDev/barter/db/database"
)

var (
	apiCfg apiConfig
)

func CreateMux(dbConn *sql.DB, secretKey string, expires_in time.Duration) *http.ServeMux {
	mux := http.NewServeMux()

	queries := database.New(dbConn)
	apiCfg.DB = queries
	apiCfg.JWTSecretKey = secretKey
	apiCfg.ExpiresIn = expires_in

	mux.HandleFunc("/users/register", apiCfg.handleRegisterUser)
	mux.HandleFunc("/users/login", apiCfg.handleLoginUser)

	return mux
}
