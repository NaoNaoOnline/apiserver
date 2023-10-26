package listhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/list"
)

// TODO
func (h *Handler) Search(ctx context.Context, req *list.SearchI) (*list.SearchO, error) {
	fmt.Printf("/list.API/Search not implemented\n")
	return &list.SearchO{}, nil
}
