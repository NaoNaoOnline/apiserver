package rulehandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
)

// TODO
func (h *Handler) Search(ctx context.Context, req *rule.SearchI) (*rule.SearchO, error) {
	fmt.Printf("/rule.API/Search not implemented\n")
	return &rule.SearchO{}, nil
}
