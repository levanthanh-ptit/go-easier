package ez_http

import (
	"net/http"

	"github.com/levanthanh-ptit/go-ez/internal/ez_http"
)

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	ez_http.JSONError(w, err, code)
}
