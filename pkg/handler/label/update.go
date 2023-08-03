package label

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/label"
)

func (h *Handler) Update(ctx context.Context, req *label.UpdateI) (*label.UpdateO, error) {
	fmt.Printf("/label.API/Update not implemented\n")
	return &label.UpdateO{}, nil
}
