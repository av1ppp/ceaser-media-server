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
	// Load config file
	var err error

	if conf, err = config.New("config.yaml"); err != nil {
		logrus.Fatal(err)
	}

	loglevel, err := logrus.ParseLevel(conf.Log.Level)
	if err != nil {
		logrus.Fatal(err)
	}

	// Logger configuration
	logrus.SetLevel(loglevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		PadLevelText: true,
	})

	// Load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	store, err := minpq.New()
	if err != nil {
		logrus.Fatal(err)
	}
	_ = store

	serv := http.NewServer(conf)
	logrus.Fatal(serv.ListenAndServe())

	// myfm, err := minio.New("ceaser-media-server", false)
	// if err != nil {
	// 	panic(err)
	// }
}
