package click

import (
	"strings"

	"github.com/golang/freetype/truetype"

	"github.com/admpub/go-captcha-assets/bindata/chars"
	//"github.com/admpub/go-captcha-assets/resources/fonts/fzshengsksjw"
	//"github.com/admpub/go-captcha-assets/resources/fonts/yrdzst"
	"github.com/admpub/go-captcha-assets/resources/images"
	"github.com/coscms/captcha/fonts"

	"github.com/admpub/go-captcha/v2/base/option"
	"github.com/admpub/go-captcha/v2/click"
)

func (a *Click) initBasic() error {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 4, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 4}),
		//click.WithRangeLen(option.RangeVal{Min: 2, Max: 4}),
		//click.WithDisabledRangeVerifyLen(true),
		click.WithRangeThumbColors([]string{
			"#1f55c4",
			"#780592",
			"#2f6b00",
			"#910000",
			"#864401",
			"#675901",
			"#016e5c",
		}),
		click.WithRangeColors([]string{
			"#fde98e",
			"#60c1ff",
			"#fcb08e",
			"#fb88ff",
			"#b4fed4",
			"#cbfaa9",
			"#78d6f8",
		}),
	)

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		return err
	}

	// fonts
	var font *truetype.Font
	if a.font != nil {
		font = a.font
	} else {
		// font, err = yrdzst.GetFont()
		// font, err = fzshengsksjw.GetFont()
		font, err = fonts.GetFont()
		if err != nil {
			return err
		}
	}

	var masterResource click.Resource
	if a.isChinese {
		masterResource = click.WithChars(chars.GetChineseChars())
	} else {
		//masterResource=click.WithChars(chars.GetAlphaChars())
		masterResource = click.WithChars(strings.Split("abcdefghijkmnpqrstuvwxy3456789ABCDEFGHJKLMNPQRSTUVWXY", ""))
	}
	// set resources
	builder.SetResources(
		masterResource,
		//click.WithChars([]string{
		//	"1A",
		//	"5E",
		//	"3d",
		//	"0p",
		//	"78",
		//	"DL",
		//	"CB",
		//	"9M",
		//}),
		//click.WithChars(chars.GetAlphaChars()),
		click.WithFonts([]*truetype.Font{font}),
		click.WithBackgrounds(imgs),
		//click.WithThumbBackgrounds(thumbImages),
	)
	a.b = builder.Make()

	// ============================

	builder.Clear()
	builder.SetOptions(
		click.WithRangeLen(option.RangeVal{Min: 4, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 4}),
		click.WithRangeThumbColors([]string{
			"#4a85fb",
			"#d93ffb",
			"#56be01",
			"#ee2b2b",
			"#cd6904",
			"#b49b03",
			"#01ad90",
		}),
	)
	builder.SetResources(
		masterResource,
		click.WithFonts([]*truetype.Font{font}),
		click.WithBackgrounds(imgs),
	)
	a.bLight = builder.Make()
	return err
}
