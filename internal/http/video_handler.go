package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/av1ppp/ceaser-media-server/internal/video"
	"github.com/go-chi/chi/v5"
)

func (s *Server) handleGetVideo(w http.ResponseWriter, r *http.Request) {
	videoID, err := strconv.Atoi(chi.URLParam(r, "videoID"))
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	v, err := s.store.Video().Get(videoID)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	// TODO: Вынести в models.go
	data, err := json.Marshal(struct {
		Title   string `json:"title"`
		FileURL string `json:"fileUrl"`
	}{
		Title:   v.Title,
		FileURL: s.conf.Server.Addr + "/files/" + v.Filename,
	})

	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
	sendStatus(w, http.StatusOK)
}

func (s *Server) handleAddVideo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		sendError(w, http.StatusBadRequest, ErrTitleNotFound)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	video := video.Video{
		Title: title,
	}

	if _, err := io.Copy(&video, file); err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	if err := s.store.Video().Save(&video); err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendStatus(w, http.StatusAccepted)
}
