package rotate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coscms/captcha"
	"github.com/wenlng/go-captcha/v2/rotate"
)

func init() {
	captcha.Register(`rotate`, NewRotate)
}

const (
	TypeBasic = `basic`
)

func NewRotate(captchaType string, store captcha.Storer) (captcha.Driver, error) {
	a := &Rotate{
		Base:   NewBase(store),
		maxAge: captcha.MaxAge,
		cType:  captchaType,
	}
	err := a.initBasic()
	return a, err
}

type Rotate struct {
	*Base
	maxAge int64 // seconds
	b      rotate.Captcha
	cType  string
}

func (a *Rotate) MakeData(ctx context.Context) (*captcha.Data, error) {
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
