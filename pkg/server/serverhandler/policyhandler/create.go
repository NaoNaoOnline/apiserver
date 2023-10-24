package policyhandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/policy"
)

func (h *Handler) Create(ctx context.Context, req *policy.CreateI) (*policy.CreateO, error) {
	fmt.Printf("/policy.API/Create not implemented\n")
	return &policy.CreateO{}, nil
}
