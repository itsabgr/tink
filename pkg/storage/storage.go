package storage

import (
	"github.com/dgraph-io/badger/v3"
)

type storage struct {
	db *badger.DB
}

const keyPrefix byte = 'K'
const valuePrefix byte = 'V'

func (s storage) get(tx *badger.Txn, prefix byte, key []byte) (value []byte, err error) {
	item, err := tx.Get(append([]byte{prefix}, key...))
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	if item.IsDeletedOrExpired() {
		return nil, ErrNotFound
	}
	return item.ValueCopy(nil)
}
func (s storage) GetByKey(key []byte) (value []byte, err error) {
	tx := s.db.NewTransaction(true)
	defer tx.Discard()
	return s.get(tx, keyPrefix, key)
}

func (s storage) GetByValue(value []byte) (key []byte, err error) {
	tx := s.db.NewTransaction(true)
	defer tx.Discard()
	return s.get(tx, valuePrefix, value)
}
func (s storage) exists(tx *badger.Txn, prefix byte, key []byte) bool {
	item, err := tx.Get(append([]byte{prefix}, key...))
	if err == nil && !item.IsDeletedOrExpired() {
		return true
	}
	return false
}
func (s storage) Add(key, value []byte) error {
	tx := s.db.NewTransaction(true)
	defer tx.Discard()
	if s.exists(tx, keyPrefix, key) || s.exists(tx, valuePrefix, value) {
		return ErrExists
	}
	err := tx.Set(append([]byte{keyPrefix}, key...), value)
	if err != nil {
		return err
	}
	err = tx.Set(append([]byte{valuePrefix}, value...), key)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s storage) Close() error {
	return s.db.Close()
}

func Open(path string) (Storage, error) {
	opt := badger.DefaultOptions(path).WithLogger(nil)
	if len(path) == 0 {
		opt.InMemory = true
	}
	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}
	return &storage{db: db}, nil
}
