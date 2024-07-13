package imgparse

import (
	"encoding/xml"
	"io"
	"strconv"
	"strings"
)

func parseWidthOrHeight(v string) (int, error) {
	v = strings.TrimSpace(v)
	v = strings.TrimSuffix(v, `px`)
	return strconv.Atoi(v)
}

func parseSVG(r io.Reader) (width int, height int, err error) {
	t := xml.NewDecoder(r)
	var to xml.Token

	for i := 0; i <= 3; i++ {
		to, err = t.Token()
		if err != nil {
			if err == io.EOF || err.Error() == "EOF" {
				err = nil
				break
			}
			return
		}

		switch v := to.(type) {
		case xml.StartElement:
			if strings.EqualFold(v.Name.Local, `svg`) {
				var found int
				for _, attr := range v.Attr {
					if strings.EqualFold(attr.Name.Local, `width`) {
						width, _ = parseWidthOrHeight(attr.Value)
						found++
						if found > 1 {
							return
						}
						continue
					}
					if strings.EqualFold(attr.Name.Local, `height`) {
						height, _ = parseWidthOrHeight(attr.Value)
						found++
						if found > 1 {
							return
						}
						continue
					}
				}
				return
			}
		}

	}

	return
}
