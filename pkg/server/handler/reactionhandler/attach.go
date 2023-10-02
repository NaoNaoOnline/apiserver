package reactionhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/reaction"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := reaction.NewAPIServer(h, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
