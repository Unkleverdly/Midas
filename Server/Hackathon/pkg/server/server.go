package server

import "net/http"

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    "192.168.0.14:" + port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}
