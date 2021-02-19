package blurhash

import "errors"

var (
	// ErrIncorrectComponents blurHash must have between 1 and 9 components
	ErrIncorrectComponents = errors.New("blurHash must have between 1 and 9 components")
	// ErrWrongSize width and height must match the pixels array
	ErrWrongSize = errors.New("width and height must match the pixels array")
	// ErrWrongBlurhashLen the blurhash string must be at least 6 characters
	ErrWrongBlurhashLen = errors.New("the blurhash string must be at least 6 characters")
	// ErrBlurhashInvalid invalid blurhash
	ErrBlurhashInvalid = errors.New("invalid blurhash")
)
