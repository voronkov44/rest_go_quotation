package utils

import (
	"errors"
	"net/http"
	"strconv"
)

func ParseAuthor(r *http.Request) (string, error) {
	author := r.PathValue("author")
	if author == "" {
		return "", errors.New("author not specified")
	}
	return author, nil
}

func ParseID(r *http.Request) (int, error) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
