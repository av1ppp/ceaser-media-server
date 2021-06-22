package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func newHandler(server *Server) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		sendStatus(w, http.StatusOK)
	})

	// Version 1
	r.Route("/v1", func(r chi.Router) {
		r.Get("/file/{filename}", server.hGetFile) // GET /v1/file/...

		r.Route("/videos", func(r chi.Router) {
			r.Post("/", server.hAddVideo)                   // POST /v1/video
			r.Get("/", server.hGetManyVideos)               // GET /v1/video
			r.Get("/{videoID:[0-9]+}", server.hGetOneVideo) // GET /v1/video/.../
			r.Delete("/{videoID:[0-9]+}", server.hDelVideo) // DELETE /v1/video/.../
		})
	})

	return r
}

func sendStatus(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(http.StatusText(statusCode)))
}

func sendError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}
