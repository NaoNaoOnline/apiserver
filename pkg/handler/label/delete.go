package label

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/pbf/label"
)

func (h *Handler) Delete(ctx context.Context, req *label.DeleteI) (*label.DeleteO, error) {
	fmt.Printf("apigocode/pkg/pbf/label.Delete not implemented\n")
	return &label.DeleteO{}, nil
}
