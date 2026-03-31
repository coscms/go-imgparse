package imgparse

import (
	"io"

	"golang.org/x/image/bmp"
)

func parseBMP(r io.Reader) (int, int, error) {
	cfg, err := bmp.DecodeConfig(r)
	if err != nil {
		return 0, 0, err
	}

	return cfg.Width, cfg.Height, nil
}
