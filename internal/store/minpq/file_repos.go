package minpq

import (
	"io"

	"github.com/av1ppp/ceaser-media-server/internal/store"
)

type fileRepository struct {
	store *Store
}

func (r *fileRepository) OpenByName(name string) (store.File, error) {
	return r.store.fm.Open(name)
}

func (r *fileRepository) GetDataByName(name string) ([]byte, error) {
	f, err := r.store.fm.Open(name)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(f)
}
