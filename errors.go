package captcha

import "errors"

var (
	// verify
	ErrParameterRequired = errors.New("[captcha]param is empty")
	ErrIllegalKey        = errors.New("[captcha]illegal key")
	ErrInvalidResponse   = errors.New("[captcha]invalid response")

	// make data
	ErrGenerateFailed     = errors.New("[captcha]generate captcha data failed")
	ErrBase64EncodeFailed = errors.New("[captcha]encoding base64 data failed")

	// common
	ErrUnsupported = errors.New("[captcha]unsupported")
)
