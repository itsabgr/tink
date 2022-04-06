package storage

import "io"

//Storage is a 1:1 map of keys to values with
type Storage interface {
	GetByKey(key []byte) (value []byte, err error)
	GetByValue(value []byte) (key []byte, err error)
	Add(key, value []byte) error
	io.Closer
}
