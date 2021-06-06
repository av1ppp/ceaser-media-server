// Minio + PostgreSQL store

package minpq

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/av1ppp/ceaser-media-server/internal/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db              *sql.DB
	videoRepository *videoRepository
}

func New() (store.Store, error) {
	var (
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
	)

	if dbname == "" {
		dbname = "ceaser_media_server"
	}

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "5432"
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s",
		user, password, dbname, host, port)

	// Подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	store := Store{db: db}
	store.videoRepository = &videoRepository{store: &store}

	return &store, nil
}

func (s *Store) Video() store.VideoRepository {
	return s.videoRepository
}
