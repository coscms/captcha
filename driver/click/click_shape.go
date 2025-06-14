package click

import (
	"github.com/admpub/go-captcha-assets/resources/images"
	"github.com/admpub/go-captcha-assets/resources/shapes"
	"github.com/admpub/go-captcha/v2/base/option"
	"github.com/admpub/go-captcha/v2/click"
)

func (a *Click) initShape() error {
	builder := click.NewBuilder(
		click.WithRangeLen(option.RangeVal{Min: 3, Max: 6}),
		click.WithRangeVerifyLen(option.RangeVal{Min: 2, Max: 3}),
		click.WithRangeThumbBgDistort(1),
		click.WithIsThumbNonDeformAbility(true),
	)

	// shape
	// click.WithUseShapeOriginalColor(false) -> Random rewriting of graphic colors
	shapeMaps, err := shapes.GetShapes()
	if err != nil {
		return err
	}

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		return err
	}

	// set resources
	builder.SetResources(
		click.WithShapes(shapeMaps),
		click.WithBackgrounds(imgs),
	)

	a.b = builder.MakeWithShape()
	return err
}
