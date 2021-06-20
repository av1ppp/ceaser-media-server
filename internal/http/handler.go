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

	r.Get("/files/{filename}", func(w http.ResponseWriter, r *http.Request) {
		data, err := server.store.File().GetDataByName(chi.URLParam(r, "filename"))
		if err != nil {
			sendError(w, http.StatusInternalServerError, err)
			return
		}

		w.Write(data)
		sendStatus(w, http.StatusOK)
	})

	r.Route("/video", func(r chi.Router) {
		r.Post("/", server.handleAddVideo)                // POST /video
		r.Get("/{videoID:[0-9]+}", server.handleGetVideo) // GET /video
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
