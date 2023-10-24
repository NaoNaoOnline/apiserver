package policyhandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/policy"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := policy.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
