package captcha

import (
	"context"
)

type Constructor func(captchaType string, store Storer) (Driver, error)

type Driver interface {
	MakeData(ctx context.Context) (*Data, error)
	Verify(ctx context.Context, key string, response string) error
}

type Storer interface {
	Put(ctx context.Context, key string, val interface{}, timeout int64) error
	Get(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
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

const MaxAge = 1800 //seconds
