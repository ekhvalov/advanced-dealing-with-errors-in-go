package reader

import (
	"fmt"
	"io"
)

var ErrInvalidChunkSize error

func ReadByChunk(r io.Reader, chunkSize int) ([][]byte, error) {
	if r == nil {
		return nil, nil
	}

	if chunkSize <= 0 {
		return nil, fmt.Errorf("%w: %v", ErrInvalidChunkSize, chunkSize)
	}

	var result [][]byte

	for {
		chunk := make([]byte, chunkSize)
		n, err := io.ReadFull(r, chunk)

		if err == io.EOF {
			break
		}
		if err == io.ErrUnexpectedEOF {
			chunk = chunk[:n]
		} else if err != nil {
			return nil, err
		}

		result = append(result, chunk)
	}

	return result, nil
}
