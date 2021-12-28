package common

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Ungzip(input []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewBuffer(input))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}
