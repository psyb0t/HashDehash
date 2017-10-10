package db

import (
	"compress/gzip"
	"encoding/base64" // Yeah not that of a good idea - buttfuck it
	"io/ioutil"
	"os"
	"path"
)

type DB struct {
	path string
}

func NewDB(path string) *DB {
	database := &DB{}
	database.SetPath(path)

	return database
}

func (db *DB) SetPath(path string) {
	db.path = path
	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}
}

func (db *DB) Get(key string) (string, error) {
	key = base64.StdEncoding.EncodeToString([]byte(key))

	f, err := os.Open(path.Join(db.path, key))
	defer f.Close()
	if err != nil {
		return "", err
	}

	gzr, err := gzip.NewReader(f)
	defer gzr.Close()
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(gzr)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (db *DB) Set(key, value string) error {
	key = base64.StdEncoding.EncodeToString([]byte(key))

	f, err := os.OpenFile(path.Join(db.path, key),
		os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	defer f.Close()
	if err != nil {
		return err
	}

	gzw := gzip.NewWriter(f)
	gzw.Write([]byte(value))
	gzw.Close()

	return nil
}
