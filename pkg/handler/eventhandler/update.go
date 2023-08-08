package eventhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
)

func (h *Handler) Update(ctx context.Context, req *event.UpdateI) (*event.UpdateO, error) {
	fmt.Printf("/event.API/Update not implemented\n")
	return &event.UpdateO{}, nil
}
