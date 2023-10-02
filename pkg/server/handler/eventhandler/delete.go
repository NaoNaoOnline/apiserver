package eventhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
)

func (h *Handler) Delete(ctx context.Context, req *event.DeleteI) (*event.DeleteO, error) {
	fmt.Printf("/event.API/Delete not implemented\n")
	return &event.DeleteO{}, nil
}
