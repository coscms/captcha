package captcha

import "fmt"

var drivers = map[string]Constructor{}

func Register(name string, constructor Constructor) {
	drivers[name] = constructor
}

func Open(name string, cType string, store Storer) (Driver, error) {
	constructor, ok := drivers[name]
	if !ok {
		return nil, ErrUnsupported
	}
	return constructor(cType, store)
}

func Unregister(name string) {
	delete(drivers, name)
}

var instances = map[string]map[string]Driver{}

func RegisterInstance(driverName string, captchaType string, instance Driver) {
	if _, ok := instances[driverName]; !ok {
		instances[driverName] = map[string]Driver{}
	}
	instances[driverName][captchaType] = instance
}

func GetInstance(driverName string, captchaType string) (Driver, error) {
	if _, ok := instances[driverName]; !ok {
		return nil, fmt.Errorf(`%w: %s`, ErrUnsupported, driverName)
	}
	if instance, ok := instances[driverName][captchaType]; !ok {
		return nil, fmt.Errorf(`%w: %s`, ErrUnsupported, captchaType)
	} else {
		return instance, nil
	}
}
