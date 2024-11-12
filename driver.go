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
	Key   string
	Image string
	Thumb string
	Tile  *Tile
}

type Tile struct {
	Image   string
	Width   int
	Height  int
	OffsetX int
	OffsetY int
}

const MaxAge = 1800 //seconds
