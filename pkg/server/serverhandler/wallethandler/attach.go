package wallethandler

import (
	"github.com/NaoNaoOnline/apigocode/pkg/wallet"
	"github.com/gorilla/mux"
)

func (h *Handler) Attach(rtr *mux.Router, opt ...interface{}) {
	han := wallet.NewAPIServer(&wrapper{han: h}, opt...)
	rtr.PathPrefix(han.PathPrefix()).Handler(han)
}
