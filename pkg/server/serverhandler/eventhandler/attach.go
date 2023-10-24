package eventhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/event"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := event.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
