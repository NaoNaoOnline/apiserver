package listhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
)

// TODO
func (h *Handler) Delete(ctx context.Context, req *list.DeleteI) (*list.DeleteO, error) {
	fmt.Printf("/list.API/Delete not implemented\n")
	return &list.DeleteO{}, nil
}
