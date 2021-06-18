package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/av1ppp/ceaser-media-server/internal/video"
)

func (s *Server) handleAddVideo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		sendStatus(w, http.StatusBadRequest)
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

	fmt.Println(video)aklwhjliawjhil
}
