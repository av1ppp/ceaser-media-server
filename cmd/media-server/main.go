package main

import (
	"os"

	"github.com/av1ppp/ceaser-media-server/internal/fm/minio"
)

func init() {
	os.Setenv("MINIO_ACCESS_KEY", ...)
	os.Setenv("MINIO_SECRET_KEY", ...)
	os.Setenv("MINIO_ENDPOINT", ...)
}

func main() {
	myfm, err := minio.New("main-bucket", false)
	if err != nil {
		panic(err)
	}
	_ = myfm

}
