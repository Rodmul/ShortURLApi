package core

import "ShortLink/internal/genErr"

var (
	ErrNotFound   = genErr.New("not found")
	ErrRepository = genErr.New("error working with repository")
)
