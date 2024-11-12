package driver

import (
	"github.com/coscms/captcha"
	"github.com/coscms/captcha/driver/click"
	"github.com/coscms/captcha/driver/rotate"
	"github.com/coscms/captcha/driver/slide"
)

func Initialize(store captcha.Storer) error {
	clickShape, err := captcha.Open(`click`, click.TypeShape, store)
	if err != nil {
		return err
	}
	captcha.RegisterInstance(`click`, click.TypeShape, clickShape)

	clickBasic, err := captcha.Open(`click`, click.TypeBasic, store)
	if err != nil {
		return err
	}
	captcha.RegisterInstance(`click`, click.TypeBasic, clickBasic)

	rotateBasic, err := captcha.Open(`rotate`, rotate.TypeBasic, store)
	if err != nil {
		return err
	}
	captcha.RegisterInstance(`rotate`, rotate.TypeBasic, rotateBasic)

	slideBasic, err := captcha.Open(`slide`, slide.TypeBasic, store)
	if err != nil {
		return err
	}
	captcha.RegisterInstance(`slide`, slide.TypeBasic, slideBasic)

	slideRegion, err := captcha.Open(`slide`, slide.TypeRegion, store)
	if err != nil {
		return err
	}
	captcha.RegisterInstance(`slide`, slide.TypeRegion, slideRegion)
	return err
}
