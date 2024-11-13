package main

import (
	"context"
	"embed"
	"net/http"

	"github.com/admpub/cache"

	"github.com/coscms/captcha"
	"github.com/coscms/captcha/driver"
)

//go:embed native
var static embed.FS

type response struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	*captcha.Data `json:",omitempty"`
}

func main() {
	store, err := cache.NewCacher(context.Background(), `memory`, cache.Options{Interval: int(captcha.MaxAge)})
	if err != nil {
		panic(err)
	}
	err = driver.Initialize(store)
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.Handle(`/`, http.FileServer(http.FS(static)))
	mux.HandleFunc(`GET /captcha/{driver}/{type}`, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c, err := captcha.GetInstance(r.PathValue(`driver`), r.PathValue(`type`))
		if err != nil {
			captcha.JSON(w, response{Code: 1, Message: err.Error()})
			return
		}
		data, err := c.MakeData(ctx)
		if err != nil {
			captcha.JSON(w, response{Code: 1, Message: err.Error()})
			return
		}
		captcha.JSON(w, response{Code: 0, Data: data})
	})
	mux.HandleFunc(`POST /captcha/{driver}/{type}`, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		c, err := captcha.GetInstance(r.PathValue(`driver`), r.PathValue(`type`))
		if err != nil {
			captcha.JSON(w, response{Code: 1, Message: err.Error()})
			return
		}
		err = c.Verify(ctx, r.FormValue(`key`), r.FormValue(`response`))
		if err != nil {
			captcha.JSON(w, response{Code: 1, Message: err.Error()})
			return
		}
		captcha.JSON(w, response{Code: 0})
	})
	println(`listen on http://127.0.0.1:4444`)
	err = http.ListenAndServe(`:4444`, mux)
	if err != nil {
		panic(err)
	}
}
