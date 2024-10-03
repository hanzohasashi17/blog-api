package storage

import "errors"

var (
	ErrPostNotFound = errors.New("post not found")
	ErrPostExists = errors.New("post exists")
)