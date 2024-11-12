package click

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coscms/captcha"
	"github.com/wenlng/go-captcha/v2/click"
)

func init() {
	captcha.Register(`click`, NewClick)
}

const (
	TypeBasic = `basic`
	TypeShape = `shape`
)

func NewClick(captchaType string, store captcha.Storer) (captcha.Driver, error) {
	a := &Click{
		Base:   NewBase(store),
		maxAge: captcha.MaxAge,
		cType:  captchaType,
	}
	var err error
	if captchaType == `shape` {
		err = a.initShape()
	} else {
		err = a.initBasic()
	}
	return a, err
}

type Click struct {
	*Base
	maxAge int64 // seconds
	b      click.Captcha
	bLight click.Captcha
	cType  string
}

func (a *Click) MakeData(ctx context.Context) (*captcha.Data, error) {
	capt := a.b
	if a.bLight != nil {
		t, y := ctx.Value(`type`).(string)
		if y && t == `light` {
			capt = a.bLight
		}
	}
	gen, err := capt.Generate()
	if err != nil {
		return nil, err
	}
	blockData := gen.GetData()
	if blockData == nil {
		return nil, captcha.ErrGenerateFailed
	}

	masterImageBase64, err := gen.GetMasterImage().ToBase64()
	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, captcha.ErrBase64EncodeFailed,err)
	}

	thumbImageBase64, err := gen.GetThumbImage().ToBase64()
	if err != nil {
		return nil, fmt.Errorf(`%w: %v`, captcha.ErrBase64EncodeFailed,err)
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
		Thumb: thumbImageBase64,
	}, nil
}
