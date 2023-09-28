package index

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const prefix = "parsed_page_" // parsed_page_100

var (
	ErrInvalidFilename     = errors.New("invalid filename")
	ErrIndexMustBePositive = errors.New("index must be > 0")
	filenameRegex          = regexp.MustCompile(`^parsed_page_(\d+)$`)
)

func GetIndexFromFileName(fileName string) (int, error) {
	// return getIndexFromFileNameSplit(fileName)
	return getIndexFromFileNameRegex(fileName)
}

func getIndexFromFileNameRegex(fileName string) (int, error) {
	result := filenameRegex.FindStringSubmatch(fileName)
	if len(result) < 2 {
		return 0, fmt.Errorf("%w: no index in filename %q", ErrInvalidFilename, fileName)
	}
	index, err := strconv.ParseInt(result[1], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("cannot parse index as int: %w", err)
	}

	if index <= 0 {
		return 0, fmt.Errorf("%w: got %d", ErrIndexMustBePositive, index)
	}

	return int(index), nil
}

func getIndexFromFileNameSplit(fileName string) (int, error) {
	parts := strings.Split(fileName, prefix)
	if len(parts) != 2 || parts[1] == "" {
		return 0, fmt.Errorf("%w: no index in filename %q", ErrInvalidFilename, fileName)
	}

	// num := parts[1]

	index, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		return 0, fmt.Errorf("cannot parse index as int: %w", err)
	}

	if index <= 0 {
		return 0, fmt.Errorf("%w: got %d", ErrIndexMustBePositive, index)
	}

	return int(index), nil
}
