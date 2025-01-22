package slide

import (
	"github.com/admpub/go-captcha-assets/resources/images"
	"github.com/admpub/go-captcha-assets/resources/tiles"
	"github.com/admpub/go-captcha/v2/slide"
)

func (a *Slide) initBasic() error {
	builder := slide.NewBuilder(
		//slide.WithGenGraphNumber(2),
		slide.WithEnableGraphVerticalRandom(true),
	)

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		return err
	}

	graphs, err := tiles.GetTiles()
	if err != nil {
		return err
	}

	newGraphs := make([]*slide.GraphImage, len(graphs))
	for i, graph := range graphs {
		newGraphs[i] = &slide.GraphImage{
			OverlayImage: graph.OverlayImage,
			MaskImage:    graph.MaskImage,
			ShadowImage:  graph.ShadowImage,
		}
	}

	// set resources
	builder.SetResources(
		slide.WithGraphImages(newGraphs),
		slide.WithBackgrounds(imgs),
	)

	a.b = builder.Make()
	return err
}
