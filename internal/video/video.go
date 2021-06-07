package video

import (
	"io"
)

type Video struct {
	ID    int32
	Title string
	File  io.ReadSeekCloser
}
