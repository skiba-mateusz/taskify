package api

import (
	"database/sql"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
	db *sql.DB
}

func NewAPIServer(listenAddr string, db *sql.DB) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	log.Println("Server is listening on", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, router)
}