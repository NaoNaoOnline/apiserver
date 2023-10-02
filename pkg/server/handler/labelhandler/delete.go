package labelhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
)

func (h *Handler) Delete(ctx context.Context, req *label.DeleteI) (*label.DeleteO, error) {
	fmt.Printf("/label.API/Delete not implemented\n")
	return &label.DeleteO{}, nil
}
