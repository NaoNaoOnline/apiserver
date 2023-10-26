package rulehandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
)

func (h *Handler) Update(ctx context.Context, req *rule.UpdateI) (*rule.UpdateO, error) {
	fmt.Printf("/rule.API/Update not implemented\n")
	return &rule.UpdateO{}, nil
}
