package main

import "net/http"

type Config struct {
	handler *http.ServeMux
	port    string
}
