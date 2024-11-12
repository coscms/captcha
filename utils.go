package captcha

import (
	"strings"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

func Md5(b []byte) string {
	m := md5.New()
	m.Write(b)
	return hex.EncodeToString(m.Sum(nil))
}

func ParseInt64(s string) (int64,error){
	return strconv.ParseInt(strings.SplitN(s,`.`,2)[0], 10, 64)
}

func JSON(w http.ResponseWriter, data interface{}, code ...int) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=utf-8")
	h.Set("X-Content-Type-Options", "nosniff")
	if len(code) > 0 && code[0] > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(b)
}
