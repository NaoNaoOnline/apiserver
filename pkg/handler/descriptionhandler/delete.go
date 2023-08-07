package descriptionhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/description"
)

func (h *Handler) Delete(ctx context.Context, req *description.DeleteI) (*description.DeleteO, error) {
	fmt.Printf("/description.API/Delete not implemented\n")
	return &description.DeleteO{}, nil
}
