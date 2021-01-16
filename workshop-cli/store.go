package cli

import (
	"io/ioutil"
	"log"
	"os"
)

type Store interface {
	Write(content []byte)
	Read() []byte
}

type FileStore struct {
	Filename string
}

func (fs *FileStore) Read() []byte {
	content, err := ioutil.ReadFile(fs.Filename)
	if err != nil {
		return []byte("")
	}
	return content
}

func (fs *FileStore) Write(content []byte) {
	f, err := os.OpenFile(fs.Filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf(">> %+v", err)
	}
	defer f.Close()
	if _, err := f.Write(content); err != nil {
		log.Fatalf(">> %+v", err)
	}
}
