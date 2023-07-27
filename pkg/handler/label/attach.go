package label

import (
	"net/http"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
)

func (h *Handler) Attach(mux *http.ServeMux, opt ...interface{}) {
	han := label.NewAPIServer(h, opt...)
	mux.Handle(han.PathPrefix(), han)
}
