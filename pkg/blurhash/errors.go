package blurhash

import "errors"

var (
	ErrIncorrectComponents = errors.New("blurHash must have between 1 and 9 components")
	ErrWrongSize           = errors.New("width and height must match the pixels array")
	ErrWrongBlurhashLen    = errors.New("the blurhash string must be at least 6 characters")
	ErrBlurhashInvalid     = errors.New("invalid blurhash")
)
