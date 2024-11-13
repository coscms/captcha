package rotate

import (
	"github.com/admpub/go-captcha/v2/base/option"
	"github.com/admpub/go-captcha/v2/rotate"
	"github.com/wenlng/go-captcha-assets/resources/images"
)

func (a *Rotate) initBasic() error {
	builder := rotate.NewBuilder(rotate.WithRangeAnglePos([]option.RangeVal{
		{Min: 20, Max: 330},
	}))

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		return err
	}

	// set resources
	builder.SetResources(
		rotate.WithImages(imgs),
	)

	a.b = builder.Make()
	return err
}
