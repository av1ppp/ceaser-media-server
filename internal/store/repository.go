package store

import "github.com/av1ppp/ceaser-media-server/internal/video"

type VideoRepository interface {
	Save(v *video.Video)
}
