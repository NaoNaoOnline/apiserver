package votehandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/vote"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := vote.NewAPIServer(h, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
