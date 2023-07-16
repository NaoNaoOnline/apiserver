package label

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/pbf/label"
)

func (h *Handler) Search(ctx context.Context, req *label.SearchI) (*label.SearchO, error) {
	fmt.Printf("apigocode/pkg/pbf/label.Search not implemented\n")
	return &label.SearchO{}, nil
}
