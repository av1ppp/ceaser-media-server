package store

type Store interface {
	Video() VideoRepository
	File() FileRepository
}
