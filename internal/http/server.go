package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	addr    string
	handler *Handler
}

// Creating new server.
func New(addr string) *Server {
	return &Server{
		addr:    addr,
		handler: newHandler(),
	}
}

type Handler struct {
	mux *chi.Mux
}

func newHandler() *Handler {
	handler := &Handler{}

	r := chi.NewRouter()

	handler.mux = r

	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// ListenAndServe listens on the TCP network address addr and then
// calls Serve with handler to handle requests on incoming connections.
func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(s.addr, s.handler)
}
