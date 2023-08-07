package descriptionhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
)

func (h *Handler) Update(ctx context.Context, req *description.UpdateI) (*description.UpdateO, error) {
	fmt.Printf("/description.API/Update not implemented\n")
	return &description.UpdateO{}, nil
}
