package captcha

import (
	"context"
)

type Constructor func(captchaType string, store Storer, options ...Option) (Driver, error)

type Driver interface {
	OptionSetter
	MakeData(ctx context.Context) (*Data, error)
	Verify(ctx context.Context, key string, response string) error
}

type OptionSetter interface {
	SetOption(key string, value interface{})
}

type Storer interface {
	Put(ctx context.Context, key string, val interface{}, timeout int64) error
	Get(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
}

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	*Data   `json:",omitempty"`
}

func (a *APIResponse) SetSuccess(message ...string) *APIResponse {
	a.Code = 0
	if len(message) > 0 {
		a.Message = message[0]
	}
	return a
}

func (a *APIResponse) SetError(message string) *APIResponse {
	a.Code = 1
	a.Message = message
	return a
}

func (a *APIResponse) SetData(data *Data) *APIResponse {
	a.Code = 0
	a.Data = data
	return a
}

func (a *APIResponse) IsSuccess() bool {
	return a.Code == 1
}

type Data struct {
	Key   string `json:"key"`
	Image string `json:"image,omitempty"`
	Thumb string `json:"thumb,omitempty"`
	Tile  *Tile  `json:"tile,omitempty"`
}

type Tile struct {
	Image   string `json:"image"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	OffsetX int    `json:"x"`
	OffsetY int    `json:"y"`
}

type Option func(d OptionSetter)

const MaxAge = 1800 //seconds
