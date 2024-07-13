package imgparse

import (
	"bytes"
	"testing"
)

func TestParseSVG(t *testing.T) {
	svg := []byte(`<?xml version="1.0" encoding="UTF-8"?>
<svg width="480px" height="102px" viewBox="0 0 480 102" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
</svg>`)
	width, height, err := parseSVG(bytes.NewReader(svg))
	if err != nil {
		t.Errorf(`Unexpected error: %v`, err)
	}
	t.Logf(`width: %d; heigh: %d`, width, height)
}
