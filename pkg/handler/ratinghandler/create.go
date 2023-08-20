package ratinghandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/rating"
)

func (h *Handler) Create(ctx context.Context, req *rating.CreateI) (*rating.CreateO, error) {
	fmt.Printf("/rating.API/Create not implemented\n")
	return &rating.CreateO{}, nil
}
