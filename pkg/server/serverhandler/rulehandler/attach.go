package rulehandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/rule"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := rule.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
