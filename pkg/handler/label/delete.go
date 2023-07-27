package label

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
)

func (h *Handler) Delete(ctx context.Context, req *label.DeleteI) (*label.DeleteO, error) {
	fmt.Printf("apigocode/pkg/label.Delete not implemented\n")
	return &label.DeleteO{}, nil
}
