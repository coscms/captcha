package captcha

import (
	"fmt"

	"github.com/webx-top/com"
)

var drivers = map[string]Constructor{}

func Register(name string, constructor Constructor) {
	drivers[name] = constructor
}

func Open(name string, cType string, store Storer, options ...Option) (Driver, error) {
	constructor, ok := drivers[name]
	if !ok {
		return nil, ErrUnsupported
	}
	return constructor(cType, store, options...)
}

func Unregister(name string) {
	delete(drivers, name)
}

var instances = com.InitSafeMap[string, Driver]()

func RegisterInstance(driverName string, captchaType string, instance Driver) {
	instances.Set(driverName+`.`+captchaType, instance)
}

func UnregisterInstance(driverName string, captchaType string) {
	instances.Delete(driverName + `.` + captchaType)
}

func GetInstance(driverName string, captchaType string) (Driver, error) {
	instance, ok := instances.GetOk(driverName + `.` + captchaType)
	if !ok {
		return nil, fmt.Errorf(`%w: %s`, ErrUnsupported, driverName+`.`+captchaType)
	}
	return instance, nil
}

func GetInstanceOk(driverName string, captchaType string) (Driver, bool) {
	return instances.GetOk(driverName + `.` + captchaType)
}
