package labelhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := label.NewAPIServer(h, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
