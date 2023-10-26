package listhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
)

// TODO
func (h *Handler) Update(ctx context.Context, req *list.UpdateI) (*list.UpdateO, error) {
	fmt.Printf("/list.API/Update not implemented\n")
	return &list.UpdateO{}, nil
}
