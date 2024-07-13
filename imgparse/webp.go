package imgparse

import (
	"errors"
	"io"
)

var (
	ErrInvalidWebpFile           = errors.New("invalid webp file")
	ErrMalformedLosslessWebpFile = errors.New("malformed lossless webp file")
	ErrMalformedWebpFile         = errors.New("malformed webp file")
	ErrUnknownWebpType           = errors.New("unknown webp type")
)

func parseWebP(r io.Reader) (int, int, error) {
	var width int
	var height int

	buf, err := readbuf(r, 31)
	if err != nil {
		return 0, 0, err
	}

	if string(buf[:4]) != "RIFF" || string(buf[8:12]) != "WEBP" {
		return 0, 0, ErrInvalidWebpFile
	}

	switch string(buf[12:16]) {
	case "VP8L":
		if buf[20] != 0x2F {
			return 0, 0, ErrMalformedLosslessWebpFile
		}

		width = (read16le(buf[21:23]) & 0x3FFF) + 1
		height = (((read16le(buf[23:25]) << 2) | (read16le(buf[21:23]) >> 14)) & 0x3FFF) + 1
	case "VP8 ":
		if (buf[20]&1) != 0 || read24le(buf[23:27]) != 0x2A019D {
			return 0, 0, ErrMalformedWebpFile
		}

		width = read16le(buf[26:28]) & 0x3fff
		height = read16le(buf[28:30]) & 0x3fff
	case "VP8X":
		width = read24le(buf[24:28]) + 1
		height = read24le(buf[27:31]) + 1
	default:
		return 0, 0, ErrUnknownWebpType
	}

	return width, height, nil
}
