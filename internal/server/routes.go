package server

import "github.com/vinicel/listEvent/auth"

func (s *Server) InitRoutes() {
	s.Router.HandleFunc("/health", auth.Ping)
}
