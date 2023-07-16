package label

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/pbf/label"
)

func (h *Handler) Update(ctx context.Context, req *label.UpdateI) (*label.UpdateO, error) {
	fmt.Printf("apigocode/pkg/pbf/label.Update not implemented\n")
	return &label.UpdateO{}, nil
}
