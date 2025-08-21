package web

import (
	"database/sql"
	"net/http"

	"github.com/MudassirDev/barter/db/database"
)

var (
	apiCfg apiConfig
)

func CreateMux(dbConn *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	queries := database.New(dbConn)
	apiCfg.DB = queries

	return mux
}
