package wal

import (
	"os"
)

type WAL struct {
	file *os.File
}

func New(path string) (*WAL, error) {

	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return &WAL{
		file: file,
	}, nil
}

func (w *WAL) Write(record string) error {

	_, err := w.file.WriteString(record + "\n")

	if err != nil {
		return err
	}

	return w.file.Sync()
}

func (w *WAL) Close() error {
	return w.file.Close()
}