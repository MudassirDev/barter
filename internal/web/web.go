package web

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/MudassirDev/barter/db/database"
)

const (
	AUTH_KEY        string = "auth_key"
	PRODUCT_ENV     string = "production"
	DEVELOPMENT_ENV string = "development"
)

var (
	apiCfg apiConfig
)

func CreateMux(dbConn *sql.DB, secretKey string, expires_in time.Duration, env string) *http.ServeMux {
	mux := http.NewServeMux()

	queries := database.New(dbConn)
	apiCfg.DB = queries
	apiCfg.JWTSecretKey = secretKey
	apiCfg.ExpiresIn = expires_in
	apiCfg.ENV = env

	mux.HandleFunc("POST /users/register", apiCfg.handleRegisterUser)
	mux.HandleFunc("POST /users/login", apiCfg.handleLoginUser)

	mux.Handle("POST /skills/create", apiCfg.authMiddleware(apiCfg.handleCreateSkill()))

	return mux
}
