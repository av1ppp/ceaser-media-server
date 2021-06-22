package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/av1ppp/ceaser-media-server/internal/video"
	"github.com/go-chi/chi/v5"
)

func (s *Server) hDelVideo(w http.ResponseWriter, r *http.Request) {
	videoID, err := strconv.Atoi(chi.URLParam(r, "videoID"))
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.store.Video().Delete(videoID); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	sendStatus(w, http.StatusOK)
}

func (s *Server) hGetManyVideos(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) hGetOneVideo(w http.ResponseWriter, r *http.Request) {
	videoID, err := strconv.Atoi(chi.URLParam(r, "videoID"))
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	v, err := s.store.Video().GetByID(videoID)
	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	// TODO: Вынести в models.go
	data, err := json.Marshal(struct {
		Title string `json:"title"`
		File  string `json:"file"`
	}{
		Title: v.Title,
		File:  "/file/" + v.Filename,
	})

	if err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (s *Server) hAddVideo(w http.ResponseWriter, r *http.Request) {
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
