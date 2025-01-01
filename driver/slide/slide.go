package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/admpub/go-captcha/v2/slide"
	"github.com/coscms/captcha"
)

func init() {
	captcha.Register(`slide`, NewSlide)
}

const (
	TypeBasic  = `basic`
	TypeRegion = `region`
)

func NewSlide(captchaType string, store captcha.Storer, options ...captcha.Option) (captcha.Driver, error) {
	a := &Slide{
		Base:   NewBase(store),
		maxAge: captcha.MaxAge,
		cType:  captchaType,
	}
	for _, option := range options {
		option(a)
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

func (a *Slide) SetOption(key string, value interface{}) {
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

	masterImageBase64, err := gen.GetMasterImage().ToBase64()
	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, captcha.ErrBase64EncodeFailed, err)
	}

	tileImageBase64, err := gen.GetTileImage().ToBase64()
	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, captcha.ErrBase64EncodeFailed, err)
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
