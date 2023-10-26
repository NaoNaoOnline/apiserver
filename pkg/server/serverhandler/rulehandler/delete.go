package rulehandler

import (
	"context"
	"fmt"

	"github.com/NaoNaoOnline/apigocode/pkg/rule"
)

// TODO
func (h *Handler) Delete(ctx context.Context, req *rule.DeleteI) (*rule.DeleteO, error) {
	fmt.Printf("/rule.API/Delete not implemented\n")
	return &rule.DeleteO{}, nil
}
