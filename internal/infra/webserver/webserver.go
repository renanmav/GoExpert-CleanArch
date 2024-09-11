package webserver

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router   chi.Router
	Handlers map[string]http.HandlerFunc
	Port     string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:   chi.NewRouter(),
		Handlers: make(map[string]http.HandlerFunc),
		Port:     webServerPort,
	}
}

func (w *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	w.Handlers[method+":"+path] = handler
}

func (w *WebServer) Start() {
	fmt.Println("Starting web server on port:", w.Port)
	w.Router.Use(middleware.Logger)
	for pathKey, handler := range w.Handlers {
		method, path, _ := strings.Cut(pathKey, ":")
		w.Router.Method(method, path, handler)
	}
	http.ListenAndServe(w.Port, w.Router)
}
