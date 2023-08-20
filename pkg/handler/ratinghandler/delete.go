package ratinghandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/rating"
)

func (h *Handler) Delete(ctx context.Context, req *rating.DeleteI) (*rating.DeleteO, error) {
	fmt.Printf("/rating.API/Delete not implemented\n")
	return &rating.DeleteO{}, nil
}
