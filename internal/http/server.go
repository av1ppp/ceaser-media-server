package http

import (
	"net/http"

	"github.com/av1ppp/ceaser-media-server/internal/config"
	"github.com/av1ppp/ceaser-media-server/internal/store"
	"github.com/sirupsen/logrus"
)

type Server struct {
	handler http.Handler
	conf    *config.Config
	store   store.Store
}

// Creating new server.
func NewServer(conf *config.Config, s store.Store) *Server {
	server := &Server{
		conf:  conf,
		store: s,
	}

	server.handler = newHandler(server)

	return server
}

// ListenAndServe listens on the TCP network address addr and then
// calls Serve with handler to handle requests on incoming connections.
func (s *Server) ListenAndServe() error {
	logrus.Debug("starting server on " + s.conf.Server.Addr)

	return http.ListenAndServe(s.conf.Server.Addr, s.handler)
}
