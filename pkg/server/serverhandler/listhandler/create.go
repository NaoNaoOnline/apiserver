package listhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
)

// TODO
func (h *Handler) Create(ctx context.Context, req *list.CreateI) (*list.CreateO, error) {
	fmt.Printf("/list.API/Create not implemented\n")
	return &list.CreateO{}, nil
}
