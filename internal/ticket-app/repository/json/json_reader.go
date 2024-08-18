package repository

import "os"

type JsonReader struct {
	path string
}

func NewJsonReader(path string) *JsonReader {
	return &JsonReader{path: path}
}

func (r JsonReader) Read(query string) ([]byte, error) {
	return os.ReadFile(r.path)
}
