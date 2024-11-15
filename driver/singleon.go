package driver

import (
	"fmt"

	"github.com/coscms/captcha"
)

func Singleton(driverName string, captchaType string, storerConstructor func() (captcha.Storer, error)) (captcha.Driver, error) {
	instance, ok := captcha.GetInstanceOk(driverName, captchaType)
	if !ok {
		store, err := storerConstructor()
		if err != nil {
			return nil, fmt.Errorf(`%w: %v`, captcha.ErrStorerInitFailed, err)
		}
		instance, err = captcha.Open(driverName, captchaType, store)
		if err != nil {
			return nil, fmt.Errorf(`%w: %v (driver=%s,type=%s)`, captcha.ErrCaptchaInitFailed, err, driverName, captchaType)
		}
		captcha.RegisterInstance(driverName, captchaType, instance)
	}
	return instance, nil
}
