package storage

type Storage interface {
	Conn() *any
}
