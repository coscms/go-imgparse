package avif

import (
	"io"

	"github.com/coscms/go-imgparse/imgparse"
	"github.com/gen2brain/avif"
)

func init() {
	imgparse.Register("avif", parseAVIF)
}

func parseAVIF(r io.Reader) (int, int, error) {
	cfg, err := avif.DecodeConfig(r)
	if err != nil {
		return 0, 0, err
	}

	return cfg.Width, cfg.Height, nil
}
