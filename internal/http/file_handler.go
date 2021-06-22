package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *Server) hGetFile(w http.ResponseWriter, r *http.Request) {
	data, err := s.store.File().GetDataByName(chi.URLParam(r, "filename"))
	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(data)
	sendStatus(w, http.StatusOK)
}
