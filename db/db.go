package db

import (
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
	data, err := ioutil.ReadFile(path.Join(db.path, key))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (db *DB) Set(key, value string) error {
	err := ioutil.WriteFile(path.Join(db.path, key), []byte(value), 0644)
	if err != nil {
		return err
	}

	return nil
}
