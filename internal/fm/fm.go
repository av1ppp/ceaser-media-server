package fm

import (
	"io"
	"time"
)

type FileManager interface {
	WriteFile(data []byte, name string) error
	Open(name string) (File, error)
	ReadFile(name string) ([]byte, error)
	ReadDir(name string) []DirEntry
	Remove(name string) error
}

// type File interface {
// 	io.Reader
// 	io.ReaderAt
// 	io.Seeker
// 	io.Closer
// }

type File io.ReadSeekCloser

type DirEntry interface {
	// Returning base filename
	Name() string

	// Returning absulute path to file
	Abs() string

	IsDir() bool
	Size() int64
	LastModified() time.Time
}
