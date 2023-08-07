package descriptionhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/description"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := description.NewAPIServer(h, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
