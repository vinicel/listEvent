package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	return &Server{
		Router: mux.NewRouter().StrictSlash(true),
	}
}

func (s *Server) Start() *Server {
	s.InitRoutes()

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"*"})

	err := http.ListenAndServe(":8900", handlers.CORS(headersOk, originsOk, methodsOk)(s.Router))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return s
}
