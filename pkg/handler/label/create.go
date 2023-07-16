package label

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/pbf/label"
)

func (h *Handler) Create(ctx context.Context, req *label.CreateI) (*label.CreateO, error) {
	fmt.Printf("apigocode/pkg/pbf/label.Create not implemented\n")
	return &label.CreateO{}, nil
}
