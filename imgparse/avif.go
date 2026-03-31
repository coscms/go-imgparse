package imgparse

import (
	"io"

	"github.com/gen2brain/avif"
)

func parseAVIF(r io.Reader) (int, int, error) {
	cfg, err := avif.DecodeConfig(r)
	if err != nil {
		return 0, 0, err
	}

	return cfg.Width, cfg.Height, nil
}
