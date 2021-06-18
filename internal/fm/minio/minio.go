package minio

import (
	"bytes"
	"context"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/av1ppp/ceaser-media-server/internal/config"
	"github.com/av1ppp/ceaser-media-server/internal/fm"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	bucketName  string
	ctx         context.Context
	minioClient *minio.Client
}

func New(conf config.MinioConfig) (fm.FileManager, error) {
	ctx := context.Background()

	if conf.EndPoint == "" {
		conf.EndPoint = "localhost:9000"
	}

	minioClient, err := minio.New(conf.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
		Secure: conf.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	exists, err := minioClient.BucketExists(ctx, conf.BucketName)
	if err != nil {
		return nil, err
	}

	if !exists {
		if err := minioClient.MakeBucket(
			ctx, conf.BucketName, minio.MakeBucketOptions{}); err != nil {
			return nil, err
		}
	}

	return &Client{
		minioClient: minioClient,
		ctx:         ctx,
		bucketName:  conf.BucketName,
	}, nil
}

func (client *Client) WriteFile(data []byte, name string) error {
	reader := bytes.NewReader(data)

	_, err := client.minioClient.PutObject(
		client.ctx, client.bucketName, name, reader, int64(len(data)), minio.PutObjectOptions{})

	return err
}

func (client *Client) Open(name string) (fm.File, error) {
	obj, err := client.minioClient.GetObject(
		client.ctx, client.bucketName, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return obj, err
}

func (client *Client) ReadFile(name string) ([]byte, error) {
	file, err := client.Open(name)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

func (client *Client) Remove(name string) error {
	// By default minio does not give an error if the file
	// does not exist, but if before that you try to get it,
	// then if the file does not exists, an error will be
	// returned.
	if _, err := client.ReadFile(name); err != nil {
		return err
	}

	return client.minioClient.RemoveObject(
		client.ctx, client.bucketName, name, minio.RemoveObjectOptions{})
}

func (client *Client) ReadDir(name string) []fm.DirEntry {
	entries := []fm.DirEntry{}

	for object := range client.minioClient.ListObjects(
		client.ctx, client.bucketName, minio.ListObjectsOptions{
			Prefix: name,
		}) {
		entries = append(entries, newDirEntry(object.Key, object.Size, object.LastModified))
	}

	return entries
}

type dirEntry struct {
	key          string
	size         int64
	lastModified time.Time
}

func (e *dirEntry) Name() string {
	return filepath.Base(e.key)
}

func (e *dirEntry) Abs() string {
	return "/" + e.key
}

func (e *dirEntry) IsDir() bool {
	return e.key[len(e.key)-1:] == "/"
}

func (e *dirEntry) Size() int64 {
	return e.size
}

func (e *dirEntry) LastModified() time.Time {
	return e.lastModified
}

func newDirEntry(key string, size int64, lastModified time.Time) *dirEntry {
	return &dirEntry{key, size, lastModified}
}
