package handler

import (
	"net/http"
)

type Interface interface {
	Attach(mux *http.ServeMux, opt ...interface{})
}
