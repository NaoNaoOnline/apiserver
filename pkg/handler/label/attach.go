package label

import (
	"github.com/NaoNaoOnline/apigocode/pkg/pbf/label"
	"google.golang.org/grpc"
)

func (h *Handler) Attach(g *grpc.Server) {
	label.RegisterAPIServer(g, h)
}
