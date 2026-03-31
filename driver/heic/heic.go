package heic

import (
	"io"

	"github.com/coscms/go-imgparse/imgparse"
	"github.com/gen2brain/heic"
)

func init() {
	imgparse.Register("heic", parseHEIC)
}

func parseHEIC(r io.Reader) (int, int, error) {
	cfg, err := heic.DecodeConfig(r)
	if err != nil {
		return 0, 0, err
	}

	return cfg.Width, cfg.Height, nil
}
