package imgparse

import (
	"errors"
	"io"
)

var ErrUnknownImageType = errors.New("unknown image type")

func ParseRes(r io.Reader, content string) (int, int, error) {
	switch content {
	case "gif":
		return parseGIF(r)
	case "jpeg":
		return parseJPEG(r)
	case "png":
		return parsePNG(r)
	case "webp":
		return parseWebP(r)
	case "webpll":
		return parseWebP(r)
	case "svg":
		return parseSVG(r)
	default:
		return 0, 0, ErrUnknownImageType
	}
}
