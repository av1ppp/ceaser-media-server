package store

import (
	"io"

	"github.com/av1ppp/ceaser-media-server/internal/video"
)

type VideoRepository interface {
	Save(v *video.Video) error
	GetByID(videoID int) (*video.Video, error)
	Delete(videoID int) error
}

type FileRepository interface {
	OpenByName(name string) (File, error)
	GetDataByName(name string) ([]byte, error)
}

type File io.ReadSeekCloser
