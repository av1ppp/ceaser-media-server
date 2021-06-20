package main

import (
	"github.com/av1ppp/ceaser-media-server/internal/config"
	"github.com/av1ppp/ceaser-media-server/internal/http"
	"github.com/av1ppp/ceaser-media-server/internal/store/minpq"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var conf *config.Config

func init() {
	// Инициализация конфига (config.yaml + environments + .env).
	var err error

	if err = godotenv.Load(); err != nil {
		panic(err)
	}

	if conf, err = config.New("config.yaml"); err != nil {
		logrus.Fatal(err)
	}

	// Конфигурация логгера
	loglevel, err := logrus.ParseLevel(conf.Log.Level)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetLevel(loglevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		PadLevelText: true,
	})
}

func main() {
	store, err := minpq.New(conf)
	if err != nil {
		logrus.Fatal(err)
	}

	serv := http.NewServer(conf, store)
	logrus.Fatal(serv.ListenAndServe())
}
