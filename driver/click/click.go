package click

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coscms/captcha"
	"github.com/wenlng/go-captcha/v2/click"
)

func NewClick(captchaType string, store captcha.Storer) *Click {
	a := &Click{
		Base:   NewBase(store),
		maxAge: captcha.MaxAge,
		cType:  captchaType,
	}
	if captchaType == `shape` {
		a.initShape()
	} else {
		a.initBasic()
	}
	return a
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

	masterImageBase64 := gen.GetMasterImage().ToBase64()
	if len(masterImageBase64) == 0 {
		return nil, captcha.ErrBase64EncodeFailed
	}

	thumbImageBase64 := gen.GetThumbImage().ToBase64()
	if len(thumbImageBase64) == 0 {
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
		Thumb: thumbImageBase64,
	}, nil
}
