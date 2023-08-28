package votehandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/vote"
)

func (h *Handler) Update(ctx context.Context, req *vote.UpdateI) (*vote.UpdateO, error) {
	fmt.Printf("/vote.API/Update not implemented\n")
	return &vote.UpdateO{}, nil
}
