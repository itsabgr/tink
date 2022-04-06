package storage

import (
	"bytes"
	"testing"
)

func TestStorage(t *testing.T) {
	db, err := Open("")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	err = db.Add([]byte("foo"), []byte("bar"))
	if err != nil {
		t.Fatal(err)
	}
	val, err := db.GetByKey([]byte("foo"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(val, []byte("bar")) {
		t.Fatal()
	}
	key, err := db.GetByValue([]byte("bar"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(key, []byte("foo")) {
		t.FailNow()
	}
	err = db.Add([]byte("foo"), []byte("bar"))
	if err == nil {
		t.Fatal()
	}

}
