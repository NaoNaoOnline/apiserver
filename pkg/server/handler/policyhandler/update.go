package policyhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
)

func (h *Handler) Update(ctx context.Context, req *policy.UpdateI) (*policy.UpdateO, error) {
	fmt.Printf("/policy.API/Update not implemented\n")
	return &policy.UpdateO{}, nil
}
