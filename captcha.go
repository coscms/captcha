package captcha

func New() *Captcha {
	return &Captcha{}
}

type Captcha struct {
	instance Driver
}
