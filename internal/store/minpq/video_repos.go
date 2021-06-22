package minpq

import (
	"github.com/av1ppp/ceaser-media-server/internal/video"
	"github.com/av1ppp/ceaser-media-server/pkg/crypto"
)

type videoRepository struct {
	store *Store
}

func (r *videoRepository) Save(v *video.Video) error {
	data, err := v.ReadAll()
	if err != nil {
		return err
	}

	filename := crypto.AsSHA256(data) + ".mp4"

	if err := r.store.fm.WriteFile(data, filename); err != nil {
		return err
	}

	_, err = r.store.db.Exec("INSERT INTO videos (title, filename) VALUES ($1, $2)",
		v.Title, filename)

	return err
}

func (r *videoRepository) GetByID(videoID int) (*video.Video, error) {
	v := video.Video{}

	row := r.store.db.QueryRow("SELECT title, filename FROM videos WHERE id=$1", videoID)
	if err := row.Scan(&v.Title, &v.Filename); err != nil {
		return nil, err
	}

	data, err := r.store.fm.ReadFile(v.Filename)
	if err != nil {
		return nil, err
	}

	v.Write(data)

	return &v, nil
}

func (r *videoRepository) Delete(videoID int) error {
	filename := ""

	row := r.store.db.QueryRow("SELECT filename FROM videos WHERE id=$1", videoID)
	if err := row.Scan(&filename); err != nil {
		return err
	}

	if err := r.store.fm.Remove(filename); err != nil {
		return err
	}

	_, err := r.store.db.Exec("DELETE FROM videos WHERE id=$1", videoID)
	return err
}
