package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/coscms/captcha"
	"github.com/wenlng/go-captcha/v2/slide"
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
	src := strings.Split(response, ",")
	if len(src) != 2 {
		return captcha.ErrInvalidResponse
	}

	var cachedData []byte
	err := a.store.Get(ctx, key, &cachedData)
	if err != nil {
		return fmt.Errorf(`%w: %s`, captcha.ErrIllegalKey, key)
	}
	var dct *slide.Block
	if err := json.Unmarshal(cachedData, &dct); err != nil {
		return fmt.Errorf(`%w: %s`, captcha.ErrIllegalKey, key)
	}

	sx, _ := strconv.ParseFloat(src[0], 64)
	sy, _ := strconv.ParseFloat(src[1], 64)
	ok := slide.CheckPoint(int64(sx), int64(sy), int64(dct.X), int64(dct.Y), 4)
	if !ok {
		return captcha.ErrInvalidResponse
	}
	return nil
}
