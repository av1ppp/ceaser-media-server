package http

import (
	"net/http"

	"github.com/av1ppp/ceaser-media-server/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	handler *Handler
	conf    *config.Config
}

// Creating new server.
func NewServer(conf *config.Config) *Server {
	return &Server{
		handler: newHandler(),
		conf:    conf,
	}
}

type Handler struct {
	mux *chi.Mux
}

func newHandler() *Handler {
	handler := &Handler{}

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		sendStatus(w, http.StatusOK)
	})

	handler.mux = r

	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// ListenAndServe listens on the TCP network address addr and then
// calls Serve with handler to handle requests on incoming connections.
func (s *Server) ListenAndServe() error {
	logrus.Debug("starting server on " + s.conf.Server.Addr)

	return http.ListenAndServe(s.conf.Server.Addr, s.handler)
}

func sendStatus(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(http.StatusText(statusCode)))
}

func sendError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}
