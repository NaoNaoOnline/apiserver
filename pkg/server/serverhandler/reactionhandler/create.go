package reactionhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/reaction"
)

func (h *Handler) Create(ctx context.Context, req *reaction.CreateI) (*reaction.CreateO, error) {
	fmt.Printf("/reaction.API/Create not implemented\n")
	return &reaction.CreateO{}, nil
}
