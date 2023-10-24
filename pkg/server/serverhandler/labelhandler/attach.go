package labelhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := label.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
