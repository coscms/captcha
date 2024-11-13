package rotate

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/admpub/go-captcha/v2/rotate"
	"github.com/coscms/captcha"
)

func NewBase(store captcha.Storer) *Base {
	return &Base{store: store}
}

type Base struct {
	store captcha.Storer
}

func (a *Base) Verify(ctx context.Context, key string, response string) error {
	if len(key) == 0 || len(response) == 0 {
		return fmt.Errorf(`%w: %s`, captcha.ErrParameterRequired, `response or key`)
	}
	sAngle, err := captcha.ParseInt64(response)
	if err != nil {
		return fmt.Errorf(`%w: %v`, captcha.ErrInvalidResponse, err)
	}
	var cachedData []byte
	err = a.store.Get(ctx, key, &cachedData)
	if err != nil {
		return fmt.Errorf(`%w: %s`, captcha.ErrIllegalKey, key)
	}
	var dct *rotate.Block
	if err := json.Unmarshal(cachedData, &dct); err != nil {
		return fmt.Errorf(`%w: %s`, captcha.ErrIllegalKey, key)
	}

	ok := rotate.CheckAngle(sAngle, int64(dct.Angle), 2)
	if !ok {
		return captcha.ErrInvalidResponse
	}
	return nil
}
