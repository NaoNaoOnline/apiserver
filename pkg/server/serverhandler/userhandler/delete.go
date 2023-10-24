package userhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/user"
)

func (h *Handler) Delete(ctx context.Context, req *user.DeleteI) (*user.DeleteO, error) {
	fmt.Printf("/user.API/Delete not implemented\n")
	return &user.DeleteO{}, nil
}
