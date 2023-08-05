package labelhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
)

func (h *Handler) Search(ctx context.Context, req *label.SearchI) (*label.SearchO, error) {
	fmt.Printf("/label.API/Search not implemented\n")
	return &label.SearchO{}, nil
}
