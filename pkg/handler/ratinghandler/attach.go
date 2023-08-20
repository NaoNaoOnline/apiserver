package ratinghandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/rating"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := rating.NewAPIServer(h, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
