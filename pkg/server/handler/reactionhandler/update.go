package reactionhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/reaction"
)

func (h *Handler) Update(ctx context.Context, req *reaction.UpdateI) (*reaction.UpdateO, error) {
	fmt.Printf("/reaction.API/Update not implemented\n")
	return &reaction.UpdateO{}, nil
}
