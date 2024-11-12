package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coscms/captcha"
	"github.com/wenlng/go-captcha/v2/slide"
)

func init() {
	captcha.Register(`slide`, NewSlide)
}

func NewSlide(captchaType string, store captcha.Storer) (captcha.Driver, error) {
	a := &Slide{
		Base:   NewBase(store),
		maxAge: captcha.MaxAge,
		cType:  captchaType,
	}
	var err error
	if captchaType == `region` {
		err = a.initRegion()
	} else {
		err = a.initBasic()
	}
	return a, err
}

type Slide struct {
	*Base
	maxAge int64 // seconds
	b      slide.Captcha
	cType  string
}

func (a *Slide) MakeData(ctx context.Context) (*captcha.Data, error) {
	gen, err := a.b.Generate()
	if err != nil {
		return nil, err
	}
	blockData := gen.GetData()
	if blockData == nil {
		return nil, captcha.ErrGenerateFailed
	}

	masterImageBase64 := gen.GetMasterImage().ToBase64()
	if len(masterImageBase64) == 0 {
		return nil, captcha.ErrBase64EncodeFailed
	}

	tileImageBase64 := gen.GetTileImage().ToBase64()
	if len(tileImageBase64) == 0 {
		return nil, captcha.ErrBase64EncodeFailed
	}

	jsonBytes, err := json.Marshal(blockData)
	if err != nil {
		return nil, fmt.Errorf(`[captcha]%w`, err)
	}
	key := captcha.Md5(jsonBytes)
	err = a.store.Put(ctx, key, jsonBytes, a.maxAge)
	if err != nil {
		return nil, fmt.Errorf(`[captcha]%w`, err)
	}
	return &captcha.Data{
		Key:   key,
		Image: masterImageBase64,
		Tile: &captcha.Tile{
			Image:   tileImageBase64,
			Width:   blockData.Width,
			Height:  blockData.Height,
			OffsetX: blockData.TileX,
			OffsetY: blockData.TileY,
		},
	}, nil
}
