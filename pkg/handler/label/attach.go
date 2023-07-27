package label

import (
	"net/http"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
	"github.com/twitchtv/twirp"
)

func (h *Handler) Attach(mux *http.ServeMux, hoo *twirp.ServerHooks) {
	han := label.NewAPIServer(h, twirp.WithServerHooks(hoo), twirp.WithServerPathPrefix(""))
	mux.Handle(han.PathPrefix(), han)
}
