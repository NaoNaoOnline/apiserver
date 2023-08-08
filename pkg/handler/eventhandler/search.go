package eventhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/event"
)

func (h *Handler) Search(ctx context.Context, req *event.SearchI) (*event.SearchO, error) {
	fmt.Printf("/event.API/Search not implemented\n")
	return &event.SearchO{}, nil
}
