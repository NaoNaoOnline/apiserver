package label

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
)

func (h *Handler) Search(ctx context.Context, req *label.SearchI) (*label.SearchO, error) {
	fmt.Printf("apigocode/pkg/label.Search not implemented\n")
	return &label.SearchO{}, nil
}
