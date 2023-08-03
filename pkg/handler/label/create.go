package label

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
)

func (h *Handler) Create(ctx context.Context, req *label.CreateI) (*label.CreateO, error) {
	fmt.Printf("/label.API/Create not implemented\n")
	return &label.CreateO{}, nil
}
