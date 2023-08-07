package descriptionhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
)

func (h *Handler) Search(ctx context.Context, req *description.SearchI) (*description.SearchO, error) {
	fmt.Printf("/description.API/Search not implemented\n")
	return &description.SearchO{}, nil
}
