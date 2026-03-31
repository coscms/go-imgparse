package imgparse

import (
	"errors"
	"io"
)

var ErrUnknownImageType = errors.New("unknown image type")

var parsers = map[string]func(io.Reader) (int, int, error){
	"gif":    parseGIF,
	"jpeg":   parseJPEG,
	"png":    parsePNG,
	"webp":   parseWebP,
	"svg":    parseSVG,
	"webpll": parseWebP,
	"tiff":   parseTIFF,
	"bmp":    parseBMP,
}

func Register(ext string, parser func(io.Reader) (int, int, error)) {
	parsers[ext] = parser
}

func Parse(r io.Reader, ext string) (int, int, error) {
	parser, exists := parsers[ext]
	if !exists {
		return 0, 0, ErrUnknownImageType
	}
	return parser(r)
}

// Deprecated: Use Parse instead.
func ParseRes(r io.Reader, ext string) (int, int, error) {
	return Parse(r, ext)
}
