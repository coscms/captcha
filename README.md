```go
package main

import (
    "context"

    "github.com/admpub/cache"

    _ "github.com/coscms/captcha/driver/click"
    _ "github.com/coscms/captcha/driver/rotate"
    _ "github.com/coscms/captcha/driver/slide"
    "github.com/coscms/captcha"
)

func main(){
    store := cache.NewCacher(context.Background(),`memory`,cache.Option{Interval:captcha.MaxAge})
    captcha.Open(`click`,`shape`,store)
}
```