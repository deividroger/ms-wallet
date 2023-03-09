package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	handlers      map[string]http.HandlerFunc
	webServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		handlers:      map[string]http.HandlerFunc{},
		webServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.handlers[path] = handler
}

func (s *WebServer) Start() error {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.handlers {
		s.Router.Post(path, handler)
	}
	return http.ListenAndServe(s.webServerPort, s.Router)
}
