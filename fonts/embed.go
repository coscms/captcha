package fonts

import (
	"embed"
	"io"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

//go:embed ArialRoundedBold.ttf
var Font embed.FS

func GetFont() (font *truetype.Font, err error) {
	file, err := Font.Open(`ArialRoundedBold.ttf`)
	if err != nil {
		return font, err
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return font, err
	}
	font, err = freetype.ParseFont(data)
	if err != nil {
		return nil, err
	}
	return font, nil
}
