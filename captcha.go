package captcha

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
