package serverhandler

import (
	"github.com/gorilla/mux"
)

type Interface interface {
	Attach(*mux.Router, ...interface{})
}
