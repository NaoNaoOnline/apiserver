package policyhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
)

func (h *Handler) Delete(ctx context.Context, req *policy.DeleteI) (*policy.DeleteO, error) {
	fmt.Printf("/policy.API/Delete not implemented\n")
	return &policy.DeleteO{}, nil
}
