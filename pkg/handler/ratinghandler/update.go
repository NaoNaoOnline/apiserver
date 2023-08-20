package ratinghandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/rating"
)

func (h *Handler) Update(ctx context.Context, req *rating.UpdateI) (*rating.UpdateO, error) {
	fmt.Printf("/rating.API/Update not implemented\n")
	return &rating.UpdateO{}, nil
}
