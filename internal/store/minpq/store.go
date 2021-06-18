// Minio + PostgreSQL store

package minpq

import (
	"database/sql"
	"fmt"

	"github.com/av1ppp/ceaser-media-server/internal/config"
	"github.com/av1ppp/ceaser-media-server/internal/fm"
	"github.com/av1ppp/ceaser-media-server/internal/fm/minio"
	"github.com/av1ppp/ceaser-media-server/internal/store"
	_ "github.com/lib/pq"
)

type Store struct {
	fm              fm.FileManager
	db              *sql.DB
	videoRepository *videoRepository
}

func New(conf *config.Config) (store.Store, error) {
	store := Store{}

	if err := store.configureDatabase(conf.DB); err != nil {
		return nil, err
	}

	if err := store.configureFileManager(conf.Minio); err != nil {
		return nil, err
	}

	store.videoRepository = &videoRepository{store: &store}

	return &store, nil
}

// Конфигурация файлового менеджера (MinIO).
func (s *Store) configureFileManager(conf config.MinioConfig) error {
	fm_, err := minio.New(conf)
	if err != nil {
		return err
	}

	s.fm = fm_
	return nil
}

// Конфигурация базы данных (PostgreSQL).
func (s *Store) configureDatabase(conf config.DBConfig) error {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d",
		conf.User, conf.Password, conf.DBName, conf.Host, conf.Port)

	// Подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Video() store.VideoRepository {
	return s.videoRepository
}
