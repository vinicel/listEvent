package server

func (s *Server) InitRoutes() {
	s.Router.HandleFunc("/health", healthCheck)
}
