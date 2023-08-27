package reactionhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/reaction"
)

func (h *Handler) Delete(ctx context.Context, req *reaction.DeleteI) (*reaction.DeleteO, error) {
	fmt.Printf("/reaction.API/Delete not implemented\n")
	return &reaction.DeleteO{}, nil
}
