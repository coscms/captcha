package click

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/coscms/captcha"
	"github.com/wenlng/go-captcha/v2/click"
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
	var cachedData []byte
	err := a.store.Get(ctx, key, &cachedData)
	if err != nil {
		return fmt.Errorf(`%w: %s`, captcha.ErrIllegalKey, key)
	}
	var dct map[int]*click.Dot
	if err := json.Unmarshal(cachedData, &dct); err != nil {
		return fmt.Errorf(`%w: %s`, captcha.ErrIllegalKey, key)
	}

	var ok bool
	if (len(dct) * 2) == len(src) {
		for i, j := 0, len(dct); i < j; i++ {
			dot := dct[i]
			j := i * 2
			k := i*2 + 1
			sx, _ := strconv.ParseFloat(src[j], 64)
			sy, _ := strconv.ParseFloat(src[k], 64)

			ok = click.CheckPoint(int64(sx), int64(sy), int64(dot.X), int64(dot.Y), int64(dot.Width), int64(dot.Height), 0)
			if !ok {
				break
			}
		}
	}

	if !ok {
		return captcha.ErrInvalidResponse
	}
	return nil
}
