package handler

import (
	"net/http"

	"github.com/twitchtv/twirp"
)

type Interface interface {
	Attach(mux *http.ServeMux, hoo *twirp.ServerHooks)
}
