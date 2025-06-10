package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fabiokusaba/go-rest-api-tutorial/services/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	// create router
	router := mux.NewRouter()

	// register services
	userStore := users.NewUserStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	// cors configuration
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:5173"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowedHeaders:   []string{"Authorization", "Content-Type"},
	// 	AllowCredentials: true,
	// })

	// corsHandler := c.Handler(router)

	log.Println("Listening on port", s.addr)

	return http.ListenAndServe(s.addr, router)
}
