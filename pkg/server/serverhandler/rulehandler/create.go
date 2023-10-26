package rulehandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
)

// TODO
func (h *Handler) Create(ctx context.Context, req *rule.CreateI) (*rule.CreateO, error) {
	fmt.Printf("/rule.API/Create not implemented\n")
	return &rule.CreateO{}, nil
}
