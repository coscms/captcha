package slide

import (
	"github.com/admpub/go-captcha-assets/resources/images"
	"github.com/admpub/go-captcha-assets/resources/tiles"
	"github.com/admpub/go-captcha/v2/slide"
)

func (a *Slide) initRegion() error {
	builder := slide.NewBuilder(
		slide.WithGenGraphNumber(2),
		slide.WithEnableGraphVerticalRandom(true),
	)

	// background image
	imgs, err := images.GetImages()
	if err != nil {
		return err
	}

	graphs, err := tiles.GetTiles()
	if err != nil {
		return err
	}
	var newGraphs = make([]*slide.GraphImage, 0, len(graphs))
	for i := 0; i < len(graphs); i++ {
		graph := graphs[i]
		newGraphs = append(newGraphs, &slide.GraphImage{
			OverlayImage: graph.OverlayImage,
			MaskImage:    graph.MaskImage,
			ShadowImage:  graph.ShadowImage,
		})
	}

	// set resources
	builder.SetResources(
		slide.WithGraphImages(newGraphs),
		slide.WithBackgrounds(imgs),
	)

	a.b = builder.MakeWithRegion()

	return err
}
